package service

import "context"

type UserInfo struct {
	Id         int64
	Mobile     int64
	Email      string
	Password   string
	Age        int
	CreateTime string
}

// UserService describes the service.
type UserService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	//用户注册
	Register(ctx context.Context, info UserInfo) (UserInfo, error)
	//根据用户ID 获取到用户数据
	UserInfoById(ctx context.Context, id int64) (UserInfo, error)
	//根据用户手机号 获取用户数据
	UserInfoByPhone(ctx context.Context, phone int64) (UserInfo, error)
}

type basicUserService struct{}

func (b *basicUserService) Register(ctx context.Context, info UserInfo) (UserInfo, error) {
	panic("implement me")
}

func (b *basicUserService) UserInfoById(ctx context.Context, id int64) (UserInfo, error) {
	panic("implement me")
}

func (b *basicUserService) UserInfoByPhone(ctx context.Context, phone int64) (UserInfo, error) {
	panic("implement me")
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService() UserService {
	return &basicUserService{}
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware) UserService {
	var svc UserService = NewBasicUserService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
