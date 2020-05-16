FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

LABEL maintainer="Kousha Godsizad <kousha.ghodsizad@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
ARG pkg

RUN go build /app/cmd/${pkg}/main.go

EXPOSE 80

CMD ["./main"]