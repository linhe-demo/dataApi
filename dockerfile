FROM golang:latest

MAINTAINER Linhe "296110717@qq.com"

WORKDIR $GOPATH/src/dataApi
ADD . $GOPATH/src/dataApi
RUN go build .

EXPOSE 8081

ENTRYPOINT ["./dataApi"]