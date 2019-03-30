FROM golang
MAINTAINER xiaobai xiaobai@qq.com
COPY ./main  /go/src/
ENTRYPOINT ["./src/main","-debug.addr=:9080","-http-addr=:9081"]