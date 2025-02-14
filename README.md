# trapipe

traQ BOTが購読したメッセージをシェルの標準出力にパイプします。\
trapipeを使用することで、CLIをtraQに組み込むことができるようになります。\
CLIがOSに影響を及ぼす機能を持っている場合、traQから間接的に実行される可能性に注意してください。

## Install

```bash
go install github.com/ras0q/trapipe@latest
```

```dockerfile
COPY --from=ghcr.io/ras0q/trapipe /bin/trapipe /bin/trapipe
```

## Usage

```bash
$ trapipe -h
Usage: trapipe --access-token=STRING <command> [flags]

Flags:
  -h, --help                           Show context-sensitive help.
      --access-token=STRING            BOT Access Token ($TRAQ_BOT_ACCESS_TOKEN)
      --ws-origin="wss://q.trap.jp"    traQ Websocket Origin ($TRAQ_WS_ORIGIN)

Commands:
  receive --access-token=STRING [flags]
    Receive messages from traQ server (default)

  send --access-token=STRING --channel-id=STRING [flags]
    Send a message to traQ server

Run "trapipe <command> --help" for more information on a command.
```

### Use with any CLIs

```bash
TRAQ_BOT_ACCESS_TOKEN="your access token"

trapipe receive -t "{{ .Message.ChannelID }} {{ .Message.PlainText }}" |
  while read -r channel_id mention args; do
    [ "$mention" = "@BOT_AWESOME" ] \
    && my-awesome-cli $args | trapipe send --channel-id "$channel_id"
  done
```

### Use within Docker

Dockerfile

```dockerfile
FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=ghcr.io/ras0q/trapipe /bin/trapipe /bin/trapipe

SHELL ["/bin/sh", "-c"]

ENTRYPOINT trapipe receive -t "{{ .Message.ChannelID }} {{ .Message.PlainText }}" | \
  while read -r channel_id mention args; do \
    [ "$mention" = "@BOT_AWESOME" ] \
    && my-awesome-cli $args | trapipe send --channel-id "$channel_id"; \
  done
```

shell

```bash
docker build -t my-awesome-image .
docker run -e TRAQ_BOT_ACCESS_TOKEN="your access token" my-awesome-image
```
