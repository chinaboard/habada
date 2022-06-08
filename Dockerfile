FROM golang:1.18.2-alpine AS builder
WORKDIR /go/src/app
ENV GOPROXY https://goproxy.io,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache ca-certificates make bash git
COPY . .
RUN bash build-linux-amd64.sh
RUN chmod +x /go/src/app/bin/habada-linux-amd64

FROM alpine:latest

MAINTAINER chinaboard <chinaboard@gmail.com>
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache \
    ca-certificates \
    tini

# install
COPY --from=builder /go/src/app/bin/habada-linux-amd64 /sbin/habada

ENTRYPOINT ["tini", "--"]
CMD ["habada"]