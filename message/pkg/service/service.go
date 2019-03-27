package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

// MessageService describes the service.
type MessageService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SendEmailMessage(ctx context.Context, email, text, content string) (err error)
	SendPhoneMessage(ctx context.Context, phone, content string) (err error)
}

type basicMessageService struct{}

func (b *basicMessageService) SendEmailMessage(ctx context.Context, email string, text string, content string) (err error) {
	// TODO implement the business logic of SendEmailMessage
	fmt.Printf("邮件发送给%s,标题为%s,内容为%s",email,text,content)
	fmt.Println()
	return errors.New("邮箱发生错误");
}
func (b *basicMessageService) SendPhoneMessage(ctx context.Context, phone string, content string) (err error) {
	// TODO implement the business logic of SendPhoneMessage
	return err
}

// NewBasicMessageService returns a naive, stateless implementation of MessageService.
func NewBasicMessageService() MessageService {
	return &basicMessageService{}
}

// New returns a MessageService with all of the expected middleware wired in.
func New(middleware []Middleware) MessageService {
	var svc MessageService = NewBasicMessageService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
