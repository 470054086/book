package service

import (
	currpb "book/curriculum/pkg/grpc/pb"
	userpb "book/user/pkg/grpc/pb"
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

type UserInfo struct {
	Id       int64            `json:"id"`
	Phone    int64            `json:"phone"`
	Password string           `json:"password"`
	Age      int64            `json:"age"`
	Curr     []CurriculumInfo `json:"curr"`
}

type CurriculumInfo struct {
	Id     int64
	UserId int64
	Title  string
	Desc   string
}

// UserapiService describes the service.
type UserapiService interface {
	// Add your methods here
	//注册-APi
	Register(ctx context.Context, phone int64, password string, age int64) (*UserInfo, error)
	//登录Api
	Login(ctx context.Context, phone int64, password string) (*UserInfo, error)
	//获取用户信息Api
	GetUserInfo(ctx context.Context, id int64) (*UserInfo, error)
	Test(ctx context.Context) error
	TestApi(ctx context.Context) error
}

type basicUserapiService struct {
	userClientGrpc userpb.UserClient       //userGrpc客户端
	currClientGrpc currpb.CurriculumClient //currGrpc客户端

}

func (b *basicUserapiService) Register(ctx context.Context, phone int64, password string, age int64) (u0 *UserInfo, e1 error) {
	// TODO implement the business logic of Register
	user := &userpb.UserInfo{
		Phone:    phone,
		Password: password,
		Age:      age,
	}
	reply, e := b.userClientGrpc.Register(context.Background(), &userpb.RegisterRequest{
		User: user,
	})

	res := &UserInfo{
		Phone:    reply.User.Phone,
		Password: reply.User.Password,
		Age:      reply.User.Age,
	}
	return res, e
}
func (b *basicUserapiService) Login(ctx context.Context, phone int64, password string) (u0 *UserInfo, e1 error) {
	// TODO implement the business logic of Login

	reply, _ := b.userClientGrpc.UserInfoByPhone(context.Background(), &userpb.UserInfoByPhoneRequest{
		Phone: phone,
	})
	res := &UserInfo{
		Id:       reply.User.Id,
		Phone:    reply.User.Phone,
		Password: reply.User.Password,
		Age:      reply.User.Age,
	}
	if reply.User.Password == password {
		return res, nil
	}
	return res, errors.New("用户名不正确")
}

func (b *basicUserapiService) Test(ctx context.Context) (e0 error) {
	// TODO implement the business logic of Test
	fmt.Print("我来测试下是通的吗")
	return e0
}

func (b *basicUserapiService) GetUserInfo(ctx context.Context, id int64) (u0 *UserInfo, e1 error) {
	// TODO implement the business logic of GetUserInfo

	userInfoChan := make(chan *userpb.UserInfoByIdReply)
	currInfoChan := make(chan *currpb.CurriculumListReply)
	var res *userpb.UserInfoByIdReply
	var currRes *currpb.CurriculumListReply

	//同时发送两个请求
	go func(c chan *userpb.UserInfoByIdReply) {
		reply, _ := b.userClientGrpc.UserInfoById(context.Background(), &userpb.UserInfoByIdRequest{
			Id: id,
		})
		c <- reply
	}(userInfoChan)
	go func(c chan *currpb.CurriculumListReply) {
		listReply, _ := b.currClientGrpc.CurriculumList(context.Background(), &currpb.CurriculumListRequest{
			UserId: id,
		})
		c <- listReply
	}(currInfoChan)
	res = <-userInfoChan
	currRes = <-currInfoChan

	ret := &UserInfo{
		Id:       res.User.Id,
		Phone:    res.User.Phone,
		Password: res.User.Password,
		Age:      res.User.Age,
	}
	var curr = []CurriculumInfo{}
	for _, v := range currRes.Info {
		var exd = CurriculumInfo{
			Id:     v.Id,
			UserId: v.UserId,
			Title:  v.Title,
			Desc:   v.Desc,
		}
		curr = append(curr, exd)
	}
	ret.Curr = curr
	return ret, e1
}

// NewBasicUserapiService returns a naive, stateless implementation of UserapiService.
func NewBasicUserapiService() UserapiService {
	Client, err := grpc.Dial("user-service:8082", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("user rpc is error %s", err))
	}
	userGrpcClient := userpb.NewUserClient(Client)

	conn, _ := grpc.Dial("curriculum-service:6082", grpc.WithInsecure())

	currGrpcClient := currpb.NewCurriculumClient(conn)

	return &basicUserapiService{
		userClientGrpc: userGrpcClient,
		currClientGrpc: currGrpcClient,
	}
}

// New returns a UserapiService with all of the expected middleware wired in.
func New(middleware []Middleware) UserapiService {
	var svc UserapiService = NewBasicUserapiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicUserapiService) TestApi(ctx context.Context) (e0 error) {
	// TODO implement the business logic of TestApi
	return e0
}
