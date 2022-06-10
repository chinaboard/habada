FROM golang:alpine AS builder
WORKDIR /go/src/app
RUN apk add --no-cache ca-certificates make bash git
ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" skipcache
RUN git clone https://github.com/chinaboard/habada.git
WORKDIR /go/src/app/habada
RUN bash build-linux-amd64.sh
RUN ls -lh
RUN chmod +x /go/src/app/habada/bin/habada-linux-amd64

FROM alpine:latest

MAINTAINER chinaboard <chinaboard@gmail.com>
RUN apk add --no-cache \
    ca-certificates \
    tini

# install
WORKDIR /opt/habada
COPY --from=builder /go/src/app/habada/bin/habada-linux-amd64 /opt/habada/habada
COPY --from=builder /go/src/app/habada/view /opt/habada/view
COPY --from=builder /go/src/app/habada/static /opt/habada/static
RUN ls -lh
ENTRYPOINT ["tini", "--"]
CMD ["/opt/habada/habada"]