package run

import (
	"fmt"

	"github.com/shanduur/discord-runner/runner"
	"github.com/sirupsen/logrus"
)

func respondError(r *runner.Runner, err error) error {
	err = fmt.Errorf("ERROR for UUID: %s : %s", r.UUID.String(), err.Error())
	logrus.Error(err.Error())
	return err
}

func runnerCreated(r *runner.Runner) string {
	const runnerCreated = `Runner created:
	- UUID: %s
	- Status: *%s*
	- URL: ` + "`%s`" + `
`
	return fmt.Sprintf(runnerCreated, r.UUID.String(), r.Status.String(), r.RepositoryURL)
}

func downloaded(r *runner.Runner) string {
	const runnerCreated = `Downloaded repo:
	- UUID: %s
	- Status: *%s*
	- URL: ` + "`%s`" + `
`
	return fmt.Sprintf(runnerCreated, r.UUID.String(), r.Status.String(), r.RepositoryURL)
}

func configOk(r *runner.Runner) string {
	const runnerCreated = `Config was loaded succesfully:
	- UUID: %s
	- Status: *%s*
	- URL: ` + "`%s`" + `
	- Config: %+v
`
	return fmt.Sprintf(runnerCreated, r.UUID.String(), r.Status.String(), r.RepositoryURL, r.Viper.AllSettings())
}

func runFinished(r *runner.Runner) string {
	const runnerCreated = `Run finished:
	- UUID: %s
	- Status: *%s*
	- URL: ` + "`%s`" + `
	- Logs: %s
`
	return fmt.Sprintf(runnerCreated, r.UUID.String(), r.Status.String(), r.RepositoryURL, r.LogsURL)
}
