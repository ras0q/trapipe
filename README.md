# trapipe

traQ BOTが購読したメッセージをシェルの標準出力にパイプします。
trapipeを使用することで、CLIをtraQに組み込むことができるようになります。
CLIがOSに影響を及ぼす機能を持っている場合、traQから間接的に実行される可能性に注意してください。

## Install

```bash
go install github.com/ras0q/trapipe@latest
```

```bash
$ trapipe --help
Usage: main [flags]

Flags:
  -h, --help                                   Show context-sensitive help.
      --access-token=STRING                    BOT Access Token ($TRAQ_BOT_ACCESS_TOKEN)
      --ws-origin="wss://q.trap.jp"            traQ Websocket Origin ($TRAQ_WS_ORIGIN)
  -t, --template="{{ .Message.PlainText }}"    Output Template (See https://pkg.go.dev/text/template)
```

## Example Usage

```bash
TRAQ_BOT_ACCESS_TOKEN="your access token"
COMMAND="my-awesome-cli"

trapipe | grep --line-buffered "^$COMMAND" | while read -r _ args; do
    $COMMAND $args
done
```
