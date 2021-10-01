package runner

import (
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

type Status int16

// constants descirbing status of runner
const (
	PENDING Status = iota
	RUNNING
	FINISHED
	ERROR
)

func (s Status) String() string {
	switch s {
	case PENDING:
		return "PENDING"

	case RUNNING:
		return "RUNNING"

	case FINISHED:
		return "FINISHED"

	case ERROR:
		return "ERROR"

	default:
		return "UNKNOWN"
	}
}

// Runner is a structure holding configuration of runner
type Runner struct {
	Status        Status
	UUID          uuid.UUID
	TmpDir        string
	Viper         *viper.Viper
	ContainerID   string
	Client        *client.Client
	RepositoryURL string
	LogsURL       string
}

// New creates new runner
func New(repositoryURL string) (r *Runner) {
	b := uuid.New()
	r = &Runner{
		UUID:          b,
		TmpDir:        path.Join("/tmp", b.String()),
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
func (r *Runner) Download() error {
	r.Status = RUNNING

	_, err := git.PlainClone(r.TmpDir, false, &git.CloneOptions{
		URL:   r.RepositoryURL,
		Depth: 1,
	})
	if err != nil {
		return errorf(r, "unable to clone the repo: %s", err.Error())
	}

	return nil
}

// Run starts all actions connected with Runner
func (r *Runner) Run() (err error) {
	r.Status = RUNNING

	r.Client, err = client.New(30 * time.Second)
	if err != nil {
		return errorf(r, "unable to create client")
	}

	r.ContainerID, err = r.Client.PrepareContainer(r.Viper.GetString("platform"))
	if err != nil {
		return errorf(r, "unable to prepare container: %s", err.Error())
	}

	return
}

// ReadCfg loads runner action configuration from repository
func (r *Runner) ReadCfg() error {
	r.Status = RUNNING

	if err := r.Viper.ReadInConfig(); err != nil {
		return errorf(r, "unable to read config: %s", err.Error())
	}

	return nil
}

// Close shuts down the runner and frees all connected resources
func (r *Runner) Close() {
	if r.Status != ERROR {
		r.Status = FINISHED
	}

	if err := os.RemoveAll(r.TmpDir); err != nil {
		logrus.Errorf("failed to remove dir: %s", err.Error())
	}

	if r.Client != nil {
		if status, err := r.Client.KillContainer(r.ContainerID); err != nil {
			logrus.Errorf(errorf(r, "failed to kill container, got status '%s': %s", status, err.Error()).Error())
		}

		if err := r.Client.RemoveContainer(r.ContainerID); err != nil {
			logrus.Errorf(errorf(r, "failed to remove container: %s", err.Error()).Error())
		}
	}
}
