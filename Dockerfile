FROM --platform=$BUILDPLATFORM golang:latest AS builder

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

FROM scratch
COPY --from=builder /bin/trapipe /bin/trapipe
