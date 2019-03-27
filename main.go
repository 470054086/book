package main

import (
	"book/message/pkg/grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main()  {
	//测试rpc框架
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if(err != nil) {
		panic(err)
	}
	reply, e := pb.NewMessageClient(conn).SendEmailMessage(context.Background(), &pb.SendEmailMessageRequest{
		Email:   "470054086@qq.com",
		Text:    "测试的邮件",
		Content: "这个是我第一封邮件啦",
	});
	if e!= nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(reply)
	fmt.Println("我返回数据了啦")
	fmt.Println(e.Error())



}
