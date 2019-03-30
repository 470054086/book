FROM golang
MAINTAINER xiaobai xiaobai@qq.com
COPY ./main  /go/src/
ENTRYPOINT ["./src/main","-debug.addr=:7080","-http-addr=:7081","-grpc-addr=:7082"]