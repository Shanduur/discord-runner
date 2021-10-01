package runner

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
	"github.com/shanduur/discord-runner/runner/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const filename = ".discord-runner"

// constants descirbing status of runner
const (
	PENDING = iota
	RUNNING
	FINISHED
	ERROR
)

// Runner is a structure holding configuration of runner
type Runner struct {
	Status        int
	TmpDir        string
	Viper         *viper.Viper
	ContainerID   string
	Client        *client.Client
	RepositoryURL string
}

// New creates new runner
func New(repositoryURL string) (r *Runner) {
	r = &Runner{
		TmpDir:        path.Join("/tmp", uuid.NewString()),
		Status:        PENDING,
		Viper:         viper.New(),
		RepositoryURL: repositoryURL,
	}

	r.Viper.SetConfigName(filename)
	r.Viper.AddConfigPath(r.TmpDir)

	return
}

// Download fetches git repo to the temporary location
// WARNING: curently only supports HTTPS git repos
func (r Runner) Download() error {
	_, err := git.PlainClone(r.TmpDir, false, &git.CloneOptions{
		URL:   r.RepositoryURL,
		Depth: 1,
	})
	if err != nil {
		return fmt.Errorf("unable to clone the repo: %s", err.Error())
	}

	return nil
}

// Run starts all actions connected with Runner
func (r *Runner) Run() (err error) {
	r.Client, err = client.New(30 * time.Second)
	if err != nil {
		return fmt.Errorf("unable to create client")
	}

	r.ContainerID, err = r.Client.PrepareContainer(r.Viper.GetString("platform"))
	if err != nil {
		return fmt.Errorf("unable to prepare container: %s", err.Error())
	}

	return
}

// ReadCfg loads runner action configuration from repository
func (r *Runner) ReadCfg() error {
	if err := r.Viper.ReadInConfig(); err != nil {
		return fmt.Errorf("unable to read config: %s", err.Error())
	}

	return nil
}

// Close shuts down the runner and frees all connected resources
func (r *Runner) Close() {
	if err := os.RemoveAll(r.TmpDir); err != nil {
		logrus.Errorf("failed to remove dir: %s", err.Error())
	}

	if r.Client != nil {
		if status, err := r.Client.KillContainer(r.ContainerID); err != nil {
			logrus.Errorf("failed to kill container, got status '%s': %s", status, err.Error())
		}
	}
}
