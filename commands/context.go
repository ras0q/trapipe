package commands

import traqwsbot "github.com/traPtitech/traq-ws-bot"

type CLI struct {
	AccessToken string `help:"BOT Access Token" env:"TRAQ_BOT_ACCESS_TOKEN"`
	WSOrigin    string `help:"traQ Websocket Origin" default:"wss://q.trap.jp" env:"TRAQ_WS_ORIGIN"`

	Receive `cmd:"" default:"1" help:"Receive messages from traQ server (default)"`
	Send    `cmd:"" help:"Send a message to traQ server"`
}

type Context struct {
	Bot *traqwsbot.Bot
}

type Runner interface {
	Run(*Context) error
}
