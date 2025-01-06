FROM golang:1.23-alpine AS builder

RUN sed -i 's/https:\/\/dl-cdn.alpinelinux.org/http:\/\/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add --no-cache curl build-base
WORKDIR /app
COPY . .
ENV GOPROXY https://goproxy.cn,direct
RUN go build -tags=jsoniter -ldflags "-s -w" -o /go/bin/app_release


FROM alpine:3.20
ENV TZ Asia/Shanghai

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    sed -i 's/https/http/g' /etc/apk/repositories && \
    apk add alpine-conf && \
    /sbin/setup-timezone -z Asia/Shanghai && \
    apk del alpine-conf

WORKDIR /opt
RUN mkdir -p /opt/conf

COPY --from=builder /go/bin/app_release /opt/app
CMD ["/opt/app"]