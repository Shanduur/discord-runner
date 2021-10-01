package commands

import (
	"github.com/lus/dgc"
	"github.com/shanduur/discord-runner/commands/run"
	"github.com/shanduur/discord-runner/commands/secret"
)

var Commands = []*dgc.Command{
	run.Cmd, secret.Cmd,
}
