# golang:alpine with version 3.16 (last alpine version also)
FROM golang@sha256:c9a90742f5457fae80d8f9f1c9fc6acd6884c749dc6c5b11c22976973564dd4f as base

RUN apk add --no-cache --update ca-certificates tzdata git && update-ca-certificates

WORKDIR /go/src/github.com/MrTimeout/dogstore/client

COPY . .

RUN go mod download && go mod verify && \
  CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-w -s" -o /go/bin ./...

FROM alpine:3.16

RUN apk add --no-cache bash

COPY --from=base /go/bin/client /client

CMD ["/client", "-h"]
