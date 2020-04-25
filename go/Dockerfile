# ベースとなるDockerイメージ指定
# Build Container
FROM golang:1.13 as builder
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/app
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
#環境変数
ARG GOPATH=/go
ARG GOROOT=/usr/local/go
ARG POSTGRES_USER=root
ARG POSTGRES_PASSWORD=password
ARG POSTGRES_DB=go_app_db
ARG POSTGRES_DB_ENDPOINT=http://postgres:5432
#モジュール管理
ENV GO111MODULE=on
# モジュールダウンロード
RUN go mod download
# Build
RUN go build -o app main.go
# Runtime Container
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/app/app /app
#環境変数
ARG POSTGRES_USER=root
ARG POSTGRES_PASSWORD=password
ARG POSTGRES_DB=go_app_db
ARG POSTGRES_DB_ENDPOINT=http://postgres:5432
EXPOSE 8080
ENTRYPOINT ["/app"]
