package commands

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
	"strings"

	"github.com/traPtitech/traq-ws-bot/payload"
)

type Receive struct {
	Template string `help:"Output Template (See https://pkg.go.dev/text/template)" default:"{{ .Message.PlainText }}" short:"t"`
}

var _ Runner = (*Receive)(nil)

func (c *Receive) Run(ctx *Context) error {
	tmpl, err := template.New("output").Parse(c.Template)
	if err != nil {
		panic(err)
	}

	ctx.Bot.OnError(func(message string) {
		slog.Error("received ERROR message", slog.String("message", message))
	})

	ctx.Bot.OnMessageCreated(func(p *payload.MessageCreated) {
		var buffer bytes.Buffer
		if err := tmpl.Execute(&buffer, p); err != nil {
			slog.Error("execute template", slog.Any("err", err))
			return
		}

		output := buffer.String()
		if strings.Contains(output, "\n") {
			slog.Error("multiline not supported now")
			return
		}

		fmt.Println(output)
	})

	if err := ctx.Bot.Start(); err != nil {
		return err
	}

	return nil
}
