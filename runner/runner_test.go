package runner_test

import (
	"os"
	"testing"

	"github.com/shanduur/discord-runner/runner"
)

var (
	URL = "https://github.com/octocat/hello-worId.git"
)

func TestDownload(t *testing.T) {
	r := runner.New()

	if err := r.Download(URL); err != nil {
		t.Errorf("got error: %s", err.Error())
	}

	if err := os.RemoveAll(r.TmpDir); err != nil {
		t.Errorf("got error: %s", err.Error())
	}
}

func TestReadCfg(t *testing.T) {
	r := runner.New()

	if err := r.Download(URL); err != nil {
		t.Errorf("got error: %s", err.Error())
	}

	if err := r.ReadCfg(); err != nil {
		t.Errorf("got error: %s", err.Error())
	}

	if err := os.RemoveAll(r.TmpDir); err != nil {
		t.Errorf("got error: %s", err.Error())
	}
}
