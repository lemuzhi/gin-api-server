FROM golang:1.20.1

MAINTAINER "lemuzhi"

WORKDIR /go/src/gin-project-template

COPY ../../ /go/src/gin-project-template

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

EXPOSE 8888
ENTRYPOINT ./server api c
