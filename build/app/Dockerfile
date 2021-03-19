FROM golang:alpine as builder

RUN mkdir -p /slinker ~/.ssh

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    bash \
    git \
    openssh

COPY go.mod go.sum /slinker/ 

WORKDIR /slinker/

RUN go mod download 

COPY . /slinker/
WORKDIR /slinker/cmd/app

RUN go build -o /slinker/binary/slinker

FROM alpine

WORKDIR /slinker
COPY --from=builder /slinker/binary/ . 

CMD ["/slinker/slinker"]