#!/bin/sh -eu

command -v jq >/dev/null || { echo "Error: jq command not found." >&2; exit 1; }
command -v trapipe >/dev/null || { echo "Error: trapipe command not found." >&2; exit 1; }
[ -z "${TRAQ_BOT_ACCESS_TOKEN}" ] && { echo "Error: TRAQ_BOT_ACCESS_TOKEN is not set." >&2; exit 1; }
[ $# -eq 0 ] && { echo "Error: No command provided to execute." >&2; exit 1; }

exec trapipe receive -t "{{ json . }}" |
  # $PAYLOAD is a MESSAGE_CREATED event payload
  # Ref: https://bot-console.trap.jp/docs/bot/events/message
  while read -r payload; do
    channel_id=$(echo "$payload" | jq -r '.message.channelId')
    cmd=$(printf "%s %s" "$@" "'$payload'")
    sh -c "$cmd" | trapipe send --channel-id "$channel_id"
  done
