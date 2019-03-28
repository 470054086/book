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
	users := make([]model.User, 0)
	err := db.G_db.Where("id = ?",id).Find(&users)

	var userInfo = &UserInfo{
		Id :      users[0].Id,
		Phone:    users[0].Phone,
		Password: users[0].Password,
		Age:      users[0].Age,
	}
	return userInfo, err
}
func (b *basicUserService) UserInfoByPhone(ctx context.Context, phone int64) (u0 *UserInfo, e1 error) {
	users := make([]model.User, 0)
	err := db.G_db.Where("phone = ?",phone).Find(&users)

	var userInfo = &UserInfo{
		Id :      users[0].Id,
		Phone:    users[0].Phone,
		Password: users[0].Password,
		Age:      users[0].Age,
	}
	return userInfo, err
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
