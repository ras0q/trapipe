FROM --platform=$BUILDPLATFORM golang:latest@sha256:20b91eda7a9627c127c0225b0d4e8ec927b476fa4130c6760928b849d769c149 AS builder

WORKDIR /app

RUN \
    --mount=type=cache,target=/go/pkg/mod/ \
    --mount=source=go.mod,target=go.mod \
    --mount=source=go.sum,target=go.sum \
    go mod download

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN \
    --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=target=. \
    go build -ldflags="-s -w" -o /bin/trapipe /app

FROM scratch AS tag-latest
COPY --from=builder /bin/trapipe /bin/trapipe

FROM alpine@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412 AS tag-alpine
RUN apk add --no-cache ca-certificates jq
COPY --from=builder /bin/trapipe /bin/trapipe
COPY ./docker-entrypoint.sh /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]

FROM golang@sha256:20b91eda7a9627c127c0225b0d4e8ec927b476fa4130c6760928b849d769c149 AS tag-golang
RUN \
  --mount=type=cache,target=/var/lib/apt,sharing=locked \
  --mount=type=cache,target=/var/cache/apt,sharing=locked \
  apt-get update && apt-get install -y ca-certificates jq
COPY --from=builder /bin/trapipe /bin/trapipe
COPY ./docker-entrypoint.sh /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]
