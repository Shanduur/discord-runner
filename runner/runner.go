package runner

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

const filename = ".discord-runner.yaml"

const (
	PENDING = iota
	RUNNING
	FINISHED
	ERROR
)

type Runner struct {
	Status int
	TmpDir string
}

func New() (r *Runner) {
	r = &Runner{}
	r.TmpDir = path.Join("/tmp", uuid.NewString())
	r.Status = PENDING

	return
}

func (r Runner) Download(url string) error {
	_, err := git.PlainClone(r.TmpDir, false, &git.CloneOptions{
		URL:   url,
		Depth: 1,
	})
	if err != nil {
		return fmt.Errorf("unable to clone the repo: %s", err.Error())
	}

	return nil
}

func (r *Runner) Run() {

}

func (r Runner) ReadCfg() error {
	f, err := os.ReadFile(path.Join(r.TmpDir, filename))
	if err != nil {
		return fmt.Errorf("unable to open config file: %s", err.Error())
	}

	cfg := make(map[string]interface{})
	if err := yaml.Unmarshal(f, cfg); err != nil {
		return fmt.Errorf("unable to unmarshall config: %s", err.Error())
	}

	return fmt.Errorf("%+v", cfg)
}

func (r Runner) cleanup() {

}
