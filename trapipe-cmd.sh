trapipe receive -t "{{ .Message.ChannelID }} {{ .Message.PlainText }}" | grep --line-buffered "$2" | while read -r cid _ args; do "$2" "$args" | trapipe send --channel-id "$cid"; done
