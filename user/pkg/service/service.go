package service

import (
	"context"
	"book/utils/db"
	"book/model"
)

type UserInfo struct {
	Id       int64
	Phone    int64
	Password string
	Age      int64
}

// UserService describes the service.
type UserService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Register(ctx context.Context, user UserInfo) (*UserInfo, error)
	UserInfoById(ctx context.Context, id int64) (*UserInfo, error)
	UserInfoByPhone(ctx context.Context, phone int64) (*UserInfo, error)
}

type basicUserService struct{}

func (b *basicUserService) Register(ctx context.Context, user UserInfo) (u0 *UserInfo, e1 error) {
	// TODO implement the business logic of Register
	//获取数据
	var users = new(model.User)
	users.Phone=13581922913
	users.Password="a8341526"
	users.Age=15
	_, e := db.G_db.Insert(users)
	if e != nil {
		return nil,e
	}
	//返回插入的数据
	return &user, e1
}
func (b *basicUserService) UserInfoById(ctx context.Context, id int64) (u0 *UserInfo, e1 error) {
	// TODO implement the business logic of UserInfoById
	return u0, e1
}
func (b *basicUserService) UserInfoByPhone(ctx context.Context, phone int64) (u0 *UserInfo, e1 error) {
	// TODO implement the business logic of UserInfoByPhone
	return u0, e1
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
