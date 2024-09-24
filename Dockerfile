FROM golang:1.22.5-bullseye

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@v1.52.3

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN go mod tidy