package main

import (
	"github.com/alecthomas/kong"
	"github.com/ras0q/trapipe/commands"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

var cli struct {
	AccessToken string `help:"BOT Access Token" env:"TRAQ_BOT_ACCESS_TOKEN" required:""`
	WSOrigin    string `help:"traQ Websocket Origin" default:"wss://q.trap.jp" env:"TRAQ_WS_ORIGIN"`

	Receive commands.Receive `cmd:"" default:"1" help:"Receive messages from traQ server (default)"`
	Send    commands.Send    `cmd:"" help:"Send a message to traQ server"`
}

func main() {
	ctx := kong.Parse(&cli)
	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: cli.AccessToken,
		Origin:      cli.WSOrigin,
	})
	if err != nil {
		panic(err)
	}

	commandCtx := commands.Context{
		Bot: bot,
	}
	if err := ctx.Run(&commandCtx); err != nil {
		panic(err)
	}
}
