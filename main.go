package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"google.golang.org/grpc"
	userpb "book/user/pkg/grpc/pb"
	"context"
)

type User struct {
	Id       int64
	Phone    int64
	Password string
	Age      int64
}

func main()  {
	//测试rpc框架
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if(err != nil) {
		panic(err)
	}
	//reply, e := pb.NewMessageClient(conn).SendEmailMessage(context.Background(), &pb.SendEmailMessageRequest{
	//	Email:   "470054086@qq.com",
	//	Text:    "测试的邮件",
	//	Content: "这个是我第一封邮件啦",
	//});
	var userInfo = &userpb.UserInfo{
		Id:1,
		Phone:13581922913,
		Password:"a8341526",
		Age:10,
	}
	registerReply, i := userpb.NewUserClient(conn).Register(context.Background(), &userpb.RegisterRequest{
		User: userInfo,
	})
	fmt.Println(i)
	fmt.Println(registerReply)

}
