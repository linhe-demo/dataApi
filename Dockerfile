FROM golang:latest

#设置工作目录
WORKDIR $GOPATH/src

#拉取项目代码
RUN git clone https://github.com/linhe-demo/dataApi.git

#添加配置文件
ADD ./config.ini $GOPATH/src/data-api

#切换工作目录
WORKDIR $GOPATH/src/data-api

#设置代理
ENV GOPROXY https://goproxy.io

#go构建可执行文件
RUN go build .

#暴露端口
EXPOSE 9001

#最终运行docker的命令
ENTRYPOINT  ["./data-api"]