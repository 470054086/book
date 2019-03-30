FROM golang
MAINTAINER xiaobai xiaobai@qq.com
COPY ./main  /go/src/
ENTRYPOINT ["./src/main","-debug.addr=:8080","-http-addr=:8081","-grpc-addr=:8082","-mysql-address=xiaobai:123456@tcp(192.168.23.1:3306)/test?charset=utf8"]