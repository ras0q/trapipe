package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

var CLI struct {
	AccessToken string `help:"BOT Access Token" env:"TRAQ_BOT_ACCESS_TOKEN"`
	WSOrigin    string `help:"traQ Websocket Origin" default:"wss://q.trap.jp" env:"TRAQ_WS_ORIGIN"`
}

func main() {
	ctx := kong.Parse(&CLI)
	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: CLI.AccessToken,
		Origin:      CLI.WSOrigin,
	})
	if err != nil {
		panic(err)
	}

	bot.OnError(func(message string) {
		ctx.Errorf("received ERROR message: %s", message)
	})

	bot.OnMessageCreated(func(p *payload.MessageCreated) {
		// TODO: 将来的にtemplateにした方が嬉しそう
		fmt.Println(p.Message.PlainText)
	})

	if err := bot.Start(); err != nil {
		panic(err)
	}
}
