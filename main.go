package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"google.golang.org/grpc"
	curriculum "book/curriculum/pkg/grpc/pb"
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
	conn, err := grpc.Dial("localhost:6082", grpc.WithInsecure())
	if(err != nil) {
		panic(err)
	}
	//reply, e := pb.NewMessageClient(conn).SendEmailMessage(context.Background(), &pb.SendEmailMessageRequest{
	//	Email:   "470054086@qq.com",
	//	Text:    "测试的邮件",
	//	Content: "这个是我第一封邮件啦",
	//});

	registerReply, i := curriculum.NewCurriculumClient(conn).Curriculum(context.Background(), &curriculum.CurriculumRequest{
		ID:1,
	})
	fmt.Println(i)
	fmt.Println(registerReply)

}
