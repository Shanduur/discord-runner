package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"
	"github.com/shanduur/discord-runner/commands"
)

type Bot struct {
	Token   string
	Session *discordgo.Session
}

func New(token string) (b Bot) {
	b.Token = token
	return
}

func (b *Bot) Run() (err error) {
	session, err := discordgo.New("Bot " + b.Token)
	if err != nil {
		return fmt.Errorf("unable to create session: %s", err.Error())
	}

	err = session.Open()
	if err != nil {
		return fmt.Errorf("unable to open session: %s", err.Error())
	}

	router := dgc.Create(&dgc.Router{
		Prefixes: []string{
			"!",
			"example!",
		},
		IgnorePrefixCase: false,
		BotsAllowed:      false,
		PingHandler: func(ctx *dgc.Ctx) {
			ctx.RespondText("Pong!")
		},
	})

	router.RegisterDefaultHelpCommand(session, nil)

	for _, c := range commands.Commands {
		router.RegisterCmd(c)
	}

	router.Initialize(session)

	return nil
}
