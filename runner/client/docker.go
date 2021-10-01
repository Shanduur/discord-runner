package client

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

// Client structure contains all info about client
type Client struct {
	dockerClient *client.Client
	timeout      time.Duration
}

// New creates new client instance according to endpoint url from ENV with default hostconfig
func New(timeout time.Duration) (c *Client, err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		err = fmt.Errorf("unable to create client: %v", err)

		return
	}

	c = &Client{
		dockerClient: cli,
		timeout:      timeout,
	}

	return
}

func (c *Client) Close() {
	if err := c.dockerClient.Close(); err != nil {
		logrus.Errorf("failed to close docker client: %s", err.Error())
	}
}

func (c Client) GetRawClient() *client.Client {
	return c.dockerClient
}

func (c Client) PrepareContainer(image string) (id string, err error) {
	ctx := context.Background()

	reader, err := c.dockerClient.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		err = fmt.Errorf("unable to pull mock container image: %v", err)

		return
	}
	io.Copy(io.Discard, reader)

	resp, err := c.dockerClient.ContainerCreate(ctx, &container.Config{
		Image:        image,
		Tty:          true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          strslice.StrSlice{"/bin/sh"},
	}, nil, nil, nil, "")
	if err != nil {
		err = fmt.Errorf("unable to create mock container: %v", err)

		return
	}

	if err = c.dockerClient.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		err = fmt.Errorf("unable to start mock container: %v", err)

		return
	}

	id = resp.ID

	return
}

func (c Client) RemoveContainer(containerID string) (err error) {
	ctx := context.Background()

	if err = c.dockerClient.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}); err != nil {
		err = fmt.Errorf("unable to remove mock container: %v", err)

		return
	}

	return
}

// ContainerStatus gets information about container's status, Can be one of "created", "running", "paused", "restarting", "removing", "exited", or "dead"
func (c Client) ContainerStatus(containerID string) (status string, err error) {
	ctx := context.Background()

	cont, err := c.dockerClient.ContainerInspect(ctx, containerID)
	if err != nil {
		err = fmt.Errorf("error on inspect: %v", err)

		return
	}

	status = cont.State.Status

	return
}

// StopContainer stops container specified by the containerID string
func (c Client) StopContainer(containerID string) (status string, err error) {
	ctx := context.Background()

	err = c.dockerClient.ContainerStop(ctx, containerID, &c.timeout)
	if err != nil {
		err = fmt.Errorf("error on stop request: %v", err)

		return
	}

	status, err = c.ContainerStatus(containerID)
	if err != nil {
		err = fmt.Errorf("error getting container status: %v", err)

		return
	}

	return
}

// StartContainer starts container specified by the containerID string
func (c Client) StartContainer(containerID string) (status string, err error) {
	ctx := context.Background()

	err = c.dockerClient.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		err = fmt.Errorf("error on start request: %v", err)

		return
	}

	status, err = c.ContainerStatus(containerID)
	if err != nil {
		err = fmt.Errorf("error getting container status: %v", err)

		return
	}

	return
}

// KillContainer kills container specified by the containerID string
func (c Client) KillContainer(containerID string) (status string, err error) {
	ctx := context.Background()

	err = c.dockerClient.ContainerKill(ctx, containerID, "SIGKILL")
	if err != nil {
		err = fmt.Errorf("error on kill request: %v", err)

		return
	}

	status, err = c.ContainerStatus(containerID)
	if err != nil {
		err = fmt.Errorf("error getting container status: %v", err)

		return
	}

	return
}

// ListContainers lists all containers running and stopped on the machine
func (c Client) ListContainers() (conts []types.Container, err error) {
	ctx := context.Background()

	conts, err = c.dockerClient.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		err = fmt.Errorf("error on list containers request: %v", err)
		return
	}

	return
}

// AttachContainer creates a pipe to stdout and stdin of container
func (c Client) AttachContainer(containerID string) (hr types.HijackedResponse, err error) {
	ctx := context.Background()

	hr, err = c.dockerClient.ContainerAttach(ctx, containerID, types.ContainerAttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
		Stderr: true,
	})
	if err != nil {
		err = fmt.Errorf("error on container attach: %v", err)
		return
	}

	return
}
