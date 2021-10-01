package notify

import "github.com/lus/dgc"

func Error(ctx *dgc.Ctx, err error) {
	ctx.RespondText(ctx.Event.Author.Mention() + " " + err.Error())
}

func Done(ctx *dgc.Ctx) {
	ctx.RespondText(ctx.Event.Author.Mention() + " Done!")
}
