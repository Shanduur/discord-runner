package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/shanduur/discord-runner/bot"
	"github.com/sirupsen/logrus"
)

var s *discordgo.Session

func main() {
	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
		logrus.Println("Bot is shutting down!")
	}()

	b := bot.New(os.Getenv("DISCORD_TOKEN"))

	if err := b.Run(); err != nil {
		logrus.Fatalf("unable to run bot: %s", err.Error())
	}

	logrus.Println("Bot is running!")
}
