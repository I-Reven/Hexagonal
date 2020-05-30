FROM golang:1.13-alpine AS build_base
LABEL maintainer="Kousha Godsizad <kousha.ghodsizad@gmail.com>"

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

ARG PKG
WORKDIR /app

COPY ./cmd ./cmd
COPY ./src ./src
COPY ./test ./test
COPY go.mod go.sum ./
COPY .test.env .live.env ./

RUN go mod download
RUN go build /app/cmd/${PKG}/main.go


FROM alpine:3.9
RUN apk add ca-certificates

ARG PKG
WORKDIR /app

COPY ./docker/application/wait-for-it.sh /usr/bin/wait-for-it
COPY ./docker/application/crontabs /etc/crontabs/root
COPY --from=build_base /app/main ./
COPY .live.env ./

CMD wait-for-it rabbitmq:5672 -- wait-for-it mongodb:27017 -- wait-for-it redis:6379 -- wait-for-it elassandra:9042 -- app/main install ; /app/main

EXPOSE 80 81 82 83 84 85 86 87 88 89