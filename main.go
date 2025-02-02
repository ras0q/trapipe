package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/alecthomas/kong"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

var cli struct {
	AccessToken string `help:"BOT Access Token" env:"TRAQ_BOT_ACCESS_TOKEN"`
	WSOrigin    string `help:"traQ Websocket Origin" default:"wss://q.trap.jp" env:"TRAQ_WS_ORIGIN"`
	Template    string `help:"Output Template (See https://pkg.go.dev/text/template)" default:"{{ .Message.PlainText }}" short:"t"`
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

	tmpl, err := template.New("output").Parse(cli.Template)
	if err != nil {
		panic(err)
	}

	bot.OnError(func(message string) {
		ctx.Errorf("received ERROR message: %s", message)
	})

	bot.OnMessageCreated(func(p *payload.MessageCreated) {
		var output bytes.Buffer
		if err := tmpl.Execute(&output, p); err != nil {
			ctx.Errorf("execute template: %s", err.Error())
			return
		}

		fmt.Println(output.String())
	})

	if err := bot.Start(); err != nil {
		panic(err)
	}
}
