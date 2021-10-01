package commands

import (
	"github.com/lus/dgc"
	"github.com/shanduur/discord-runner/commands/run"
)

var Commands = []*dgc.Command{
	run.Cmd,
}
