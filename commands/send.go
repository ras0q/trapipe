package commands

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/traPtitech/go-traq"
)

type Send struct {
	ChannelID string `help:"Channel ID to send a message" required:""`
}

var _ Runner = (*Send)(nil)

func (c *Send) Run(ctx *Context) error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	message := string(data)
	_, resp, err := ctx.Bot.API().
		ChannelAPI.
		PostMessage(context.Background(), c.ChannelID).
		PostMessageRequest(traq.PostMessageRequest{Content: message}).
		Execute()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("invalid status (%d %s)", resp.StatusCode, resp.Status)
	}

	return nil
}
