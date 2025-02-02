# trapipe

traQ BOTが購読したメッセージをシェルの標準出力にパイプします。\
trapipeを使用することで、CLIをtraQに組み込むことができるようになります。\
CLIがOSに影響を及ぼす機能を持っている場合、traQから間接的に実行される可能性に注意してください。\

## Install

```bash
go install github.com/ras0q/trapipe@latest
```

```dockerfile
COPY --from=ghcr.io/ras0q/trapipe /bin/trapipe /bin/trapipe
```

## Usage

```bash
$ trapipe --help
Usage: main [flags]

Flags:
  -h, --help                                   Show context-sensitive help.
      --access-token=STRING                    BOT Access Token ($TRAQ_BOT_ACCESS_TOKEN)
      --ws-origin="wss://q.trap.jp"            traQ Websocket Origin ($TRAQ_WS_ORIGIN)
  -t, --template="{{ .Message.PlainText }}"    Output Template (See https://pkg.go.dev/text/template)
```

### Use with any CLIs

```bash
TRAQ_BOT_ACCESS_TOKEN="your access token"
COMMAND="my-awesome-cli"

trapipe | grep --line-buffered "^$COMMAND" | while read -r _ args; do
    $COMMAND $args
done
```

### Use within Docker

Dockerfile

```dockerfile
FROM ubuntu:latest

RUN apt update && apt install -y ca-certificates

COPY --from=ghcr.io/ras0q/trapipe /bin/trapipe /bin/trapipe

# 色々な処理...

ENTRYPOINT ["/bin/bash", "-c", "/bin/trapipe | grep --line-buffered '^my-awesome-cli' | while read -r _ args; do my-awesome-cli $args; done"]
```

shell

```bash
docker build -t my-awesome-image .
docker run -e TRAQ_BOT_ACCESS_TOKEN="your access token" my-awesome-image
```
