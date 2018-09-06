///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package functionmanager

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	dapi "github.com/vmware/dispatch/pkg/api/v1"
	"github.com/vmware/dispatch/pkg/function-manager/backend"
	fnrunner "github.com/vmware/dispatch/pkg/function-manager/gen/restapi/operations/runner"
	fnstore "github.com/vmware/dispatch/pkg/function-manager/gen/restapi/operations/store"
	"github.com/vmware/dispatch/pkg/trace"
	"github.com/vmware/dispatch/pkg/utils"
	"path/filepath"
	"os"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	FunctionType = "Function"

func kubeClientConfig(kubeconfPath string) (*rest.Config, error) {
	if kubeconfPath == "" {
		userKubeConfig := filepath.Join(os.Getenv("HOME"), ".kube/config")
		if _, err := os.Stat(userKubeConfig); err == nil {
			kubeconfPath = userKubeConfig
		}
	}
	if kubeconfPath != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfPath)
	}
	return rest.InClusterConfig()
}

func knClient(kubeconfPath string) knclientset.Interface {
	config, err := kubeClientConfig(kubeconfPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error configuring k8s API client"))
	}
	return knclientset.NewForConfigOrDie(config)
}

>>>>>>> e1196a0a... Knative Eventing PoC
type knHandlers struct {
	backend    backend.Backend
	httpClient *http.Client
}

// NewHandlers is the constructor for the function manager API knHandlers
func NewHandlers(kubeconfPath string) Handlers {
	return &knHandlers{backend: backend.Knative(kubeconfPath), httpClient: &http.Client{}}
}

func (h *knHandlers) addFunction(params fnstore.AddFunctionParams) middleware.Responder {
	span, ctx := trace.Trace(params.HTTPRequest.Context(), "")
	defer span.Finish()

	org := *params.XDispatchOrg
	project := *params.XDispatchProject

	function := params.Body
	utils.AdjustMeta(&function.Meta, dapi.Meta{Org: org, Project: project})

	createdFunction, err := h.backend.Add(ctx, function)
	if err != nil {
		if _, ok := err.(backend.AlreadyExists); ok {
			return fnstore.NewAddFunctionConflict().WithPayload(&dapi.Error{
				Code:    http.StatusConflict,
				Message: utils.ErrorMsgAlreadyExists("function", function.Meta.Name),
			})
		}
		log.Errorf("%+v", errors.Wrap(err, "creating a function"))
		return fnstore.NewAddFunctionDefault(500).WithPayload(&dapi.Error{
			Code:    http.StatusInternalServerError,
			Message: utils.ErrorMsgInternalError("function", function.Meta.Name),
		})
	}

	return fnstore.NewAddFunctionCreated().WithPayload(createdFunction)
}

func (h *knHandlers) getFunction(params fnstore.GetFunctionParams) middleware.Responder {
	span, ctx := trace.Trace(params.HTTPRequest.Context(), "")
	defer span.Finish()

	org := *params.XDispatchOrg
	project := *params.XDispatchProject

	name := params.FunctionName

	function, err := h.backend.Get(ctx, &dapi.Meta{Name: name, Org: org, Project: project})
	if err != nil {
		if _, ok := err.(backend.NotFound); ok {
			return fnstore.NewGetFunctionNotFound().WithPayload(&dapi.Error{
				Code:    http.StatusNotFound,
				Message: utils.ErrorMsgNotFound("function", name),
			})
		}
		errors.Wrapf(err, "getting function '%s'", name)
	}

	return fnstore.NewGetFunctionOK().WithPayload(function)
}

func (h *knHandlers) deleteFunction(params fnstore.DeleteFunctionParams) middleware.Responder {
	span, ctx := trace.Trace(params.HTTPRequest.Context(), "")
	defer span.Finish()

	org := *params.XDispatchOrg
	project := *params.XDispatchProject

	name := params.FunctionName

	err := h.backend.Delete(ctx, &dapi.Meta{Name: name, Org: org, Project: project})

	if err != nil {
		if _, ok := err.(backend.NotFound); ok {
			return fnstore.NewDeleteFunctionNotFound().WithPayload(&dapi.Error{
				Code:    http.StatusNotFound,
				Message: utils.ErrorMsgNotFound("function", name),
			})
		}
		errors.Wrapf(err, "deleting function '%s'", name)
	}

	return fnstore.NewDeleteFunctionOK()
}

func (h *knHandlers) getFunctions(params fnstore.GetFunctionsParams) middleware.Responder {
	span, ctx := trace.Trace(params.HTTPRequest.Context(), "")
	defer span.Finish()

	org := *params.XDispatchOrg
	project := *params.XDispatchProject

	functions, err := h.backend.List(ctx, &dapi.Meta{Org: org, Project: project})
	if err != nil {
		log.Errorf("%+v", errors.Wrap(err, "listing functions"))
		return fnstore.NewGetFunctionsDefault(500).WithPayload(&dapi.Error{
			Code:    http.StatusInternalServerError,
			Message: swag.String(err.Error()),
		})
	}

	return fnstore.NewGetFunctionsOK().WithPayload(functions)
}

func (h *knHandlers) updateFunction(params fnstore.UpdateFunctionParams) middleware.Responder {
	span, ctx := trace.Trace(params.HTTPRequest.Context(), "")
	defer span.Finish()

	org := *params.XDispatchOrg
	project := *params.XDispatchProject

	function := params.Body
	utils.AdjustMeta(&function.Meta, dapi.Meta{Org: org, Project: project})

	updatedFunction, err := h.backend.Update(ctx, function)
	if err != nil {
		if _, ok := err.(backend.NotFound); ok {
			return fnstore.NewUpdateFunctionNotFound().WithPayload(&dapi.Error{
				Code:    http.StatusNotFound,
				Message: utils.ErrorMsgNotFound("function", function.Meta.Name),
			})
		}
		log.Errorf("%+v", errors.Wrap(err, "updating a function"))
		return fnstore.NewUpdateFunctionDefault(500).WithPayload(&dapi.Error{
			Code:    http.StatusInternalServerError,
			Message: utils.ErrorMsgInternalError("function", function.Meta.Name),
		})
	}

	return fnstore.NewUpdateFunctionOK().WithPayload(updatedFunction)
}

func (h *knHandlers) runFunction(params fnrunner.RunFunctionParams) middleware.Responder {
	span, ctx := trace.Trace(params.HTTPRequest.Context(), "")
	defer span.Finish()

	org := *params.XDispatchOrg
	project := *params.XDispatchProject
	name := *params.FunctionName

	contentType := params.Body.HTTPContext["Content-Type"].(string)
	accept := params.Body.HTTPContext["Accept"].(string)
	inBytes := params.Body.InputBytes

	runEndpoint, err := h.backend.RunEndpoint(ctx, &dapi.Meta{Name: name, Org: org, Project: project})
	if err != nil {
		if _, ok := err.(backend.NotFound); ok {
			return fnrunner.NewRunFunctionNotFound().WithPayload(&dapi.Error{
				Code:    http.StatusNotFound,
				Message: utils.ErrorMsgNotFound("function", name),
			})
		}
		errors.Wrapf(err, "getting function '%s'", name)
	}

	req, err := http.NewRequest("POST", runEndpoint, bytes.NewReader(inBytes))
	if err != nil {
		log.Errorf("%+v", errors.Wrap(err, "building http request"))
		return fnrunner.NewRunFunctionDefault(500).WithPayload(&dapi.Error{
			Code:    http.StatusInternalServerError,
			Message: utils.ErrorMsgInternalError("building http request to run function", name),
		})
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)

	response, err := h.httpClient.Do(req)
	if err != nil {
		log.Errorf("%+v", errors.Wrap(err, "performing http request"))
		return fnrunner.NewRunFunctionDefault(502).WithPayload(&dapi.Error{
			Code:    http.StatusBadGateway,
			Message: utils.ErrorMsgInternalError("performing http request to run function", name),
		})
	}
	defer response.Body.Close()

	outContentType := response.Header.Get("Content-Type")
	outBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("%+v", errors.Wrap(err, "reading http response body"))
		return fnrunner.NewRunFunctionDefault(502).WithPayload(&dapi.Error{
			Code:    http.StatusBadGateway,
			Message: utils.ErrorMsgInternalError("reading http response body running function", name),
		})
	}

	run := &dapi.Run{
		HTTPContext: map[string]interface{}{"Content-Type": outContentType},
		OutputBytes: outBytes,
	}
	return fnrunner.NewRunFunctionOK().WithPayload(run)
}

func (*knHandlers) getRun(params fnrunner.GetRunParams) middleware.Responder {
	panic("implement me")
}

func (*knHandlers) getRuns(params fnrunner.GetRunsParams) middleware.Responder {
	panic("implement me")
}
