FROM golang:1.18-alpine

RUN apk update \
  && apk add --no-cache \
    build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
RUN go install ./cmd/CaptainJack

RUN which CaptainJack

CMD ["CaptainJack"]