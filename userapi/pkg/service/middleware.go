package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserapiService) UserapiService

type loggingMiddleware struct {
	logger log.Logger
	next   UserapiService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserapiService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserapiService) UserapiService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Register(ctx context.Context, phone int64, password string, age int64) (u0 *UserInfo, e1 error) {
	defer func() {
		l.logger.Log("method", "Register", "phone", phone, "password", password, "age", age, "u0", u0, "e1", e1)
	}()
	return l.next.Register(ctx, phone, password, age)
}
func (l loggingMiddleware) Login(ctx context.Context, phone int64, password string) (u0 *UserInfo, e1 error) {
	defer func() {
		l.logger.Log("method", "Login", "phone", phone, "password", password, "u0", u0, "e1", e1)
	}()
	return l.next.Login(ctx, phone, password)
}
func (l loggingMiddleware) GetUserInfo(ctx context.Context, id int64) (u0 *UserInfo, e1 error) {
	defer func() {
		l.logger.Log("method", "GetUserInfo", "id", id, "u0", u0, "e1", e1)
	}()
	return l.next.GetUserInfo(ctx, id)
}

func (l loggingMiddleware) Test(ctx context.Context) (e1 error) {
	defer func() {
		l.logger.Log("method", "Test", "e1", e1)
	}()
	return l.next.Test(ctx)
}

func (l loggingMiddleware) TestApi(ctx context.Context) (e0 error) {
	defer func() {
		l.logger.Log("method", "TestApi", "e0", e0)
	}()
	return l.next.TestApi(ctx)
}
