FROM golang:latest

MAINTAINER xueqi
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/go
COPY . $GOPATH/src/go
RUN go build .

EXPOSE 8080
CMD ./xueqiya