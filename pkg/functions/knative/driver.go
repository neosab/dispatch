///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package knative

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	servingApi "github.com/knative/serving/pkg/apis/serving/v1alpha1"
	serving "github.com/knative/serving/pkg/client/clientset/versioned"
	servingClientV1Alpha1 "github.com/knative/serving/pkg/client/clientset/versioned/typed/serving/v1alpha1"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/vmware/dispatch/pkg/functions"
	"github.com/vmware/dispatch/pkg/trace"
	"github.com/vmware/dispatch/pkg/utils"
)

const (
	jsonContentType = "application/json"

	defaultCreateTimeout = 60 // seconds
)

// Config contains the Knative configuration
type Config struct {
	K8sConfig       string
	FuncNamespace   string
	CreateTimeout   *int
	ImagePullSecret string
}

type knativeDriver struct {
	servingClient servingClientV1Alpha1.ServingV1alpha1Interface
	fnNs          string

	createTimeout   int
	imagePullSecret string
}

type systemError struct {
	Err error `json:"err"`
}

func (err *systemError) Error() string {
	return err.Err.Error()
}

func (err *systemError) AsSystemErrorObject() interface{} {
	return err
}

func (err *systemError) StackTrace() errors.StackTrace {
	if e, ok := err.Err.(functions.StackTracer); ok {
		return e.StackTrace()
	}

	return nil
}

// New creates a new Knative driver
func New(config *Config) (functions.FaaSDriver, error) {
	k8sConf, err := kubeClientConfig(config.K8sConfig)
	servingClient, err := serving.NewForConfig(k8sConf)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error configuring k8s API client"))
	}

	fnNs := config.FuncNamespace
	if fnNs == "" {
		fnNs = "default"
	}

	d := &knativeDriver{
		servingClient: servingClient.ServingV1alpha1(),
		fnNs:          fnNs,
		createTimeout: defaultCreateTimeout,
	}
	if config.CreateTimeout != nil {
		d.createTimeout = *config.CreateTimeout
	}
	if config.ImagePullSecret != "" {
		d.imagePullSecret = config.ImagePullSecret
	}

	return d, nil
}

func kubeClientConfig(kubeConfPath string) (*rest.Config, error) {
	if kubeConfPath != "" {
		return clientcmd.BuildConfigFromFlags("", kubeConfPath)
	}
	return rest.InClusterConfig()
}

func getID(id string) string {
	return fmt.Sprintf("kn-%s", id)
}

func (d *knativeDriver) Create(ctx context.Context, f *functions.Function) error {
	span, ctx := trace.Trace(ctx, "")
	defer span.Finish()

	knService := &servingApi.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: getID(f.FaasID),
		},
		Spec: servingApi.ServiceSpec{
			RunLatest: &servingApi.RunLatestType{
				Configuration: servingApi.ConfigurationSpec{
					RevisionTemplate: servingApi.RevisionTemplateSpec{
						Spec: servingApi.RevisionSpec{
							Container: corev1.Container{
								Image: f.FunctionImageURL,
							},
						},
					},
				},
			},
		},
	}
	_, err := d.servingClient.Services(d.fnNs).Create(knService)
	if err != nil {
		return err
	}

	// make sure the function has started
	return utils.Backoff(time.Duration(d.createTimeout)*time.Second, func() error {
		knService, err := d.servingClient.Services(d.fnNs).Get(getID(f.FaasID), metav1.GetOptions{})
		if err != nil {
			return errors.Wrapf(err, "failed to read function deployment status: '%s'", f.Name)
		}

		for _, condition := range knService.Status.Conditions {
			if condition.Type == servingApi.ServiceConditionRoutesReady && condition.Status == corev1.ConditionTrue {
				return nil
			}
		}

		return errors.Errorf("function deployment not available: '%s'", f.Name)
	})
}

func (d *knativeDriver) Delete(ctx context.Context, f *functions.Function) error {
	span, ctx := trace.Trace(ctx, "")
	defer span.Finish()
	err := d.servingClient.Services(d.fnNs).Delete(getID(f.FaasID), &metav1.DeleteOptions{})
	if err != nil && !k8sErrors.IsNotFound(err) {
		return err
	}
	return nil
}

func (d *knativeDriver) doHTTPReq(faasID string, body []byte) ([]byte, error) {
	// TODO: Add istio namespace in configuration
	req, err := http.NewRequest("POST", "http://knative-ingressgateway.istio-system.svc.cluster.local", bytes.NewReader(body))

	if err != nil {
		return nil, fmt.Errorf("Unable to create request %v", err)
	}
	// TODO: Get the internal domain from knative service
	req.Host = fmt.Sprintf("%s.%s.svc.cluster.local", getID(faasID), d.fnNs)
	req.Header.Add("Content-Type", jsonContentType)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error: received error code %d: %s", resp.StatusCode, resp.Status)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *knativeDriver) GetRunnable(e *functions.FunctionExecution) functions.Runnable {
	return func(ctx functions.Context, in interface{}) (interface{}, error) {
		bytesIn, _ := json.Marshal(functions.Message{Context: ctx, Payload: in})
		res, err := d.doHTTPReq(e.FaasID, bytesIn)
		if err != nil {
			return nil, err
		}
		var out functions.Message
		if err := json.Unmarshal(res, &out); err != nil {
			return nil, &systemError{errors.Errorf("cannot JSON-parse result from Knative: %s %s", err, string(res))}
		}
		ctx.AddLogs(out.Context.Logs())
		ctx.SetError(out.Context.GetError())
		return out.Payload, nil
	}
}
