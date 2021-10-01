package run

import (
	"time"

	"github.com/lus/dgc"
)

var Cmd = &dgc.Command{
	Name:        "run",
	Description: "responds with the echo of arg",
	Usage:       "run [Git URL]",
	Example:     "run git@github.com:octocat/hello-worId.git",

	Flags: []string{},

	IgnoreCase:  false,
	SubCommands: []*dgc.Command{},

	RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
		ctx.RespondText("You are being rate limited!")
	}),

	Handler: handler,
}

func handler(ctx *dgc.Ctx) {
	// TODO:
	// - generate uuid for the action
	//   - start subroutine to registry
	// - download git repo to TMP folder
	//   - notify user about finished download
	// - run CI according to .discord-runner.yaml
	//   - process should report current progress and status
	//   - on failure stop
	// - report final status
	//   - upload logs somewhere and send formatted message to channel

	if ctx.Arguments.Amount() != 1 {
		ctx.RespondText("Wrong number of arguments!")
		return
	}

	ctx.RespondText("Got: " + ctx.Arguments.Get(0).Raw())
}
