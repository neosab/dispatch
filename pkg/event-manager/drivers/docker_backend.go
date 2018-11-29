///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package drivers

// NO TEST

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	docker "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/vmware/dispatch/pkg/client"
	"github.com/vmware/dispatch/pkg/entity-store"
	"github.com/vmware/dispatch/pkg/event-manager/drivers/entities"
	"github.com/vmware/dispatch/pkg/events/driverclient"
	"github.com/vmware/dispatch/pkg/utils"
)

const (
	labelEventDriverID   = "dispatch-eventdriver-id"
	defaultDeployTimeout = 10 // seconds
)

// DockerClient specifies the Docker client API interface required by docker driver
type DockerClient interface {
	docker.ContainerAPIClient
	docker.ImageAPIClient
}

type dockerBackend struct {
	dockerClient  DockerClient
	secretsClient client.SecretsClient
	eventsGateway string
	DeployTimeout int
}

// NewDockerBackend creates a new docker backend driver
func NewDockerBackend(dockerClient DockerClient, secretsClient client.SecretsClient, eventsGateway string) (Backend, error) {
	return &dockerBackend{
		dockerClient:  dockerClient,
		secretsClient: secretsClient,
		DeployTimeout: defaultDeployTimeout,
		eventsGateway: eventsGateway,
	}, nil
}

func bindEnv(secrets map[string]string) []string {
	var vars []string
	for key, val := range secrets {
		// ENV=value
		envVar := strings.Replace(strings.ToUpper(key), "-", "_", -1) + "=" + val
		vars = append(vars, envVar)
	}
	return vars
}

func (d *dockerBackend) getDriverSecrets(ctx context.Context, driver *entities.Driver) (map[string]string, error) {
	secrets := make(map[string]string)
	for _, secretName := range driver.Secrets {
		secret, err := d.secretsClient.GetSecret(ctx, driver.OrganizationID, secretName)
		if err != nil {
			return nil, err
		}
		for key, val := range secret.Secrets {
			secrets[key] = val
		}
	}
	return secrets, nil
}

// Expose Already combine deploy and expose in one function, throw not implement error
func (d *dockerBackend) Expose(ctx context.Context, driver *entities.Driver) error {
	return errors.New("Not Implement Error")
}

// Deploy event driver
func (d *dockerBackend) Deploy(ctx context.Context, driver *entities.Driver) error {
	log.Infof("Docker backend: exposing driver %v", driver.Name)
	secrets, err := d.getDriverSecrets(ctx, driver)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("err getting secrets of driver %s", driver.Image))
	}

	rc, err := d.dockerClient.ImagePull(ctx, driver.Image, types.ImagePullOptions{})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("err pulling image %s", driver.Image))
	}
	defer rc.Close()
	io.Copy(os.Stdout, rc)

	secrets[driverclient.AuthToken] = driver.ID
	flags := buildArgs(driver.Config)
	if d.eventsGateway != "" {
		flags = append(flags, fmt.Sprintf("--%s=%s", driverclient.DispatchEventsGatewayFlag, d.eventsGateway))
	}

	return utils.Backoff(time.Duration(d.DeployTimeout)*time.Second, func() error {
		config := &container.Config{
			Image: driver.Image,
			Env:   bindEnv(secrets),
			Cmd:   flags,
			Labels: map[string]string{
				labelEventDriverID: driver.ID,
			},
		}

		hostConfig := &container.HostConfig{}

		if driver.Expose {
			config.ExposedPorts = nat.PortSet{
				"80/tcp": struct{}{},
			}
			hostConfig.PortBindings = nat.PortMap{
				"80/tcp": []nat.PortBinding{
					{
						HostIP:   "0.0.0.0",
						HostPort: "0",
					},
				},
			}
		}

		created, err := d.dockerClient.ContainerCreate(ctx, config, hostConfig, nil, "")
		if err != nil {
			return errors.Wrap(err, "error creating container")
		}

		containerID := created.ID

		if err := d.dockerClient.ContainerStart(ctx, created.ID, types.ContainerStartOptions{}); err != nil {
			return errors.Wrap(err, "error starting container")
		}
		if driver.Expose {
			cDetails, err := d.dockerClient.ContainerInspect(ctx, containerID)
			if err != nil {
				return errors.Wrapf(err, "error when inspecting container with ID %s", containerID)
			}
			binding, ok := cDetails.NetworkSettings.Ports["80/tcp"]
			if !ok || len(binding) < 1 {
				return errors.Errorf("No port assigned to eventdriver container, docker error or no more ports available")
			}
			driver.URL = fmt.Sprintf("http://0.0.0.0:%s", binding[0].HostPort)
		}

		return nil
	})
}

// Update updates driver
func (d *dockerBackend) Update(ctx context.Context, driver *entities.Driver) error {
	// Update the driver in UPDATING status
	// docker doesn't support change env/args of running container,
	// will first delete then create
	if driver.Status == entitystore.StatusUPDATING {
		if err := d.Delete(ctx, driver); err != nil {
			return errors.Wrap(err, "error deleting container before update")
		}
		if err := d.Deploy(ctx, driver); err != nil {
			return errors.Wrap(err, "error deploying container during update")
		}
		log.Infof("driver %s updated", driver.Name)
	}
	return nil
}

// Delete deletes driver
func (d *dockerBackend) Delete(ctx context.Context, driver *entities.Driver) error {
	log.Infof("Docker backend: deleting driver %v", driver.Name)

	opts := filters.NewArgs()
	opts.Add("label", labelEventDriverID+"="+driver.ID)
	containers, err := d.dockerClient.ContainerList(ctx, types.ContainerListOptions{
		Filters: opts,
		All:     true,
	})
	if len(containers) != 1 || err != nil {
		return errors.Wrap(err, "error getting container while deleting event driver")
	}

	err = d.dockerClient.ContainerRemove(ctx, containers[0].ID, types.ContainerRemoveOptions{
		Force: true,
	})

	return err
}

func buildArgs(input map[string]string) []string {
	var args []string
	for key, val := range input {
		if val == "" {
			args = append(args, fmt.Sprintf("--%s", key))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", key, val))
		}

	}
	return args
}