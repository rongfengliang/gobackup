FROM golang:1.15-alpine AS build-env
WORKDIR /go/src/app
ENV  GO111MODULE=on
ENV  GOPROXY=https://goproxy.cn
COPY . .
RUN apk update && apk add git \
    && go build

FROM alpine:latest
RUN set -x \
    && /bin/sed -i 's,http://dl-cdn.alpinelinux.org,https://mirrors.aliyun.com,g' /etc/apk/repositories \
    && apk update && apk add ca-certificates mongodb-tools mysql-client redis postgresql-client && rm -rf /var/cache/apk/*
COPY --from=build-env /go/src/app/gobackup /usr/bin/gobackup
COPY gobackup.yml /etc/gobackup/gobackup.yml
ENTRYPOINT [ "gobackup" ]
