package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/traPtitech/traq-ws-bot/payload"
	"golang.org/x/sync/errgroup"
)

type Receive struct {
	Template string `help:"Output Template (See https://pkg.go.dev/text/template)" default:"{{ .Message.PlainText }}" short:"t"`
}

var _ Runner = (*Receive)(nil)

func (c *Receive) Run(ctx *Context) error {
	tmpl, err := template.New("output").
		Funcs(template.FuncMap{
			"json": func(v any) string {
				encoded, _ := json.Marshal(v)
				return string(encoded)
			},
		}).
		Parse(c.Template)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	errCh := make(chan error, 1)

	ctx.Bot.OnError(func(message string) {
		errCh <- fmt.Errorf("received ERROR message: %s", message)
	})

	ctx.Bot.OnMessageCreated(func(p *payload.MessageCreated) {
		var buffer bytes.Buffer
		if err := tmpl.Execute(&buffer, p); err != nil {
			errCh <- fmt.Errorf("execute template: %w", err)
			return
		}

		output := buffer.String()
		if strings.Contains(output, "\n") {
			errCh <- fmt.Errorf("multiline is not supported. Use `json` function")
			return
		}

		fmt.Println(output)
	})

	eg := errgroup.Group{}
	eg.Go(ctx.Bot.Start)
	eg.Go(func() error {
		if err := <-errCh; err != nil {
			return fmt.Errorf("handler error: %w", err)
		}

		return nil
	})

	return eg.Wait()
}
