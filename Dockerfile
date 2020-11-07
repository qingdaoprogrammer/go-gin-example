FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/github.com/han/go-gin-example
COPY . $GOPATH/github.com/han/go-gin-example
RUN go build .

EXPOSE 8009
ENTRYPOINT ["./go-gin-example"]