FROM golang:1.11.8 as builder
ADD . /go/src/github.com/crosscode-nl/incwebdemo
WORKDIR /go/src/github.com/crosscode-nl/incwebdemo
RUN go get github.com/gomodule/redigo/redis
RUN CGO_ENABLED=0 go build -o incwebdemo

FROM alpine
LABEL maintainer="Patrick Vollebregt <patrick@crosscode.nl>"
EXPOSE 5000
ENV REDISADDR redis:6379
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /go/src/github.com/crosscode-nl/incwebdemo/incwebdemo /app/
WORKDIR /app
CMD ["./incwebdemo"]
