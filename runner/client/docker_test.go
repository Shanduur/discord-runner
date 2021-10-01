package client_test

import (
	"testing"
	"time"

	"github.com/shanduur/discord-runner/runner/client"
)

const IMAGE = "alpine"

func TestNewClient_Smoke(t *testing.T) {
	_, err := client.New(30 * time.Second)
	if err != nil {
		t.Errorf("error creating client: %v", err)
	}

}

func TestClientMethods(t *testing.T) {
	var (
		err         error
		containerID string
		status      string
		c           *client.Client
	)

	if c, err = client.New(10 * time.Second); err != nil {
		t.Errorf("error during creating client: %v", err)
	}
	defer c.Close()

	if containerID, err = c.PrepareContainer(IMAGE); err != nil {
		t.Errorf("error during preparing test: %v", err)
	}

	if status, err = c.StopContainer(containerID); err != nil {
		t.Errorf("error during stoping container: %v", err)
	} else if status != "exited" {
		t.Errorf("stoping container: got: %v wanted :%v", status, "exited")
	}

	if status, err = c.StartContainer(containerID); err != nil {
		t.Errorf("error during starting container: %v", err)
	} else if status != "running" {
		t.Errorf("starting container: got: %v wanted :%v", status, "running")
	}

	if status, err = c.KillContainer(containerID); err != nil {
		t.Errorf("error during killing container: %v", err)
	} else if status != "exited" {
		t.Errorf("killing container: got: %v wanted :%v", status, "exited")
	}

	if err = c.RemoveContainer(containerID); err != nil {
		t.Errorf("error during performing cleanup %v", err)
	}
}

func TestListContainers(t *testing.T) {
	var (
		err error
		c   *client.Client
	)

	if c, err = client.New(10 * time.Second); err != nil {
		t.Errorf("error during creating client: %v", err)
	}
	defer c.Close()

	if _, err = c.ListContainers(); err != nil {
		t.Errorf("error during getting list of container: %v", err)
	}
}
