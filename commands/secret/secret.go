package secret

import (
	"fmt"
	"regexp"
	"time"

	"github.com/lus/dgc"
	"github.com/shanduur/discord-runner/notify"
)

var Cmd = &dgc.Command{
	Name:        "secret",
	Description: "list, add or remove secrets",
	Usage:       "secret (ls|add|rm) [KEY](=VALUE)",
	Example:     "secret add KEY=VALUE",

	Flags: []string{},

	IgnoreCase:  false,
	SubCommands: []*dgc.Command{},

	RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
		ctx.RespondText("You are being rate limited!")
	}),

	Handler: handler,
}

func handler(ctx *dgc.Ctx) {
	if ctx.Arguments.Amount() == 0 {
		ctx.RespondText("Wrong number of arguments!")
		return
	}

	switch ctx.Arguments.Get(0).Raw() {
	case "ls":
		kvs, err := ls()
		if err != nil {
			notify.Error(ctx, respondError(err))
			return
		}
		ctx.RespondText(table(kvs))
	case "add":
		if err := add(ctx.Arguments.Get(1).Raw()); err != nil {
			notify.Error(ctx, respondError(err))
			return
		}
		ctx.RespondText("OK")
	case "rm":
		if err := rm(ctx.Arguments.Get(1).Raw()); err != nil {
			notify.Error(ctx, respondError(err))
			return
		}
		ctx.RespondText("OK")
	}
}

func ls() (kvs map[string]string, err error) {
	kvs = make(map[string]string)
	kvs["KEY-1"] = "VALUE-1"
	kvs["KEY-2"] = "VALUE-2"
	kvs["KEY-3"] = "VALUE-3"

	return
}

func rm(key string) error {
	return nil
}

func add(keyValue string) error {
	expr := `(\w+)=([^\s]+)`
	reg := regexp.MustCompile(expr)

	if !reg.MatchString(keyValue) {
		return fmt.Errorf("Key/Value pair is incorrect. Should match regular expression: %s", expr)
	}

	return nil
}
