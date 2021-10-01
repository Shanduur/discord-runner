package runner_test

import (
	"testing"

	"github.com/shanduur/discord-runner/runner"
)

var (
	URL = "https://github.com/Shanduur/discord-runner.git"
)

func TestDownload(t *testing.T) {
	r := runner.New(URL)
	defer r.Close()

	if err := r.Download(); err != nil {
		t.Errorf("got error: %s", err.Error())
	}
}

func TestReadCfg(t *testing.T) {
	r := runner.New(URL)
	defer r.Close()

	if err := r.Download(); err != nil {
		t.Errorf("got error: %s", err.Error())
	}

	if err := r.ReadCfg(); err != nil {
		t.Errorf("got error: %s", err.Error())
	}
}
