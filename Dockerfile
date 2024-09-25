FROM golang:1.22.5-bullseye

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@v1.52.3

RUN go install github.com/a-h/templ/cmd/templ@latest

# 这里必须要指定 -tags，不然不会下载依赖，执行 migrate 命令迁移是会报错（nknown driver mysql (forgotten import?)）
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY . .

RUN go mod tidy