package endpoint

import (
	"book/userapi/pkg/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

// RegisterRequest collects the request parameters for the Register method.
type RegisterRequest struct {
	Phone    int64  `json:"phone"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
}

// RegisterResponse collects the response parameters for the Register method.
type RegisterResponse struct {
	U0 *service.UserInfo `json:"u0"`
	E1 error             `json:"e1"`
}

// MakeRegisterEndpoint returns an endpoint that invokes Register on the service.
func MakeRegisterEndpoint(s service.UserapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		u0, e1 := s.Register(ctx, req.Phone, req.Password, req.Age)
		return RegisterResponse{
			E1: e1,
			U0: u0,
		}, nil
	}
}

// Failed implements Failer.
func (r RegisterResponse) Failed() error {
	return r.E1
}

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Phone    int64  `json:"phone"`
	Password string `json:"password"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	U0 *service.UserInfo `json:"u0"`
	E1 error             `json:"e1"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.UserapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		u0, e1 := s.Login(ctx, req.Phone, req.Password)
		return LoginResponse{
			E1: e1,
			U0: u0,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.E1
}

// GetUserInfoRequest collects the request parameters for the GetUserInfo method.
type GetUserInfoRequest struct {
	Id int64 `json:"id"`
}

// GetUserInfoResponse collects the response parameters for the GetUserInfo method.
type GetUserInfoResponse struct {
	U0 *service.UserInfo `json:"u0"`
	E1 error             `json:"e1"`
}

// MakeGetUserInfoEndpoint returns an endpoint that invokes GetUserInfo on the service.
func MakeGetUserInfoEndpoint(s service.UserapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserInfoRequest)
		u0, e1 := s.GetUserInfo(ctx, req.Id)
		return GetUserInfoResponse{
			E1: e1,
			U0: u0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserInfoResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Register implements Service. Primarily useful in a client.
func (e Endpoints) Register(ctx context.Context, phone int64, password string, age int64) (u0 *service.UserInfo, e1 error) {
	request := RegisterRequest{
		Age:      age,
		Password: password,
		Phone:    phone,
	}
	response, err := e.RegisterEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RegisterResponse).U0, response.(RegisterResponse).E1
}

// Login implements Service. Primarily useful in a client.
func (e Endpoints) Login(ctx context.Context, phone int64, password string) (u0 *service.UserInfo, e1 error) {
	request := LoginRequest{
		Password: password,
		Phone:    phone,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).U0, response.(LoginResponse).E1
}

// GetUserInfo implements Service. Primarily useful in a client.
func (e Endpoints) GetUserInfo(ctx context.Context, id int64) (u0 *service.UserInfo, e1 error) {
	request := GetUserInfoRequest{Id: id}
	response, err := e.GetUserInfoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserInfoResponse).U0, response.(GetUserInfoResponse).E1
}

// TestRequest collects the request parameters for the Test method.
type TestRequest struct{}

// TestResponse collects the response parameters for the Test method.
type TestResponse struct {
	E0 error `json:"e0"`
}

// MakeTestEndpoint returns an endpoint that invokes Test on the service.
func MakeTestEndpoint(s service.UserapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		e0 := s.Test(ctx)
		return TestResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r TestResponse) Failed() error {
	return r.E0
}

// Test implements Service. Primarily useful in a client.
func (e Endpoints) Test(ctx context.Context) (e0 error) {
	request := TestRequest{}
	response, err := e.TestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(TestResponse).E0
}

// TestApiRequest collects the request parameters for the TestApi method.
type TestApiRequest struct{}

// TestApiResponse collects the response parameters for the TestApi method.
type TestApiResponse struct {
	E0 error `json:"e0"`
}

// MakeTestApiEndpoint returns an endpoint that invokes TestApi on the service.
func MakeTestApiEndpoint(s service.UserapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		e0 := s.TestApi(ctx)
		return TestApiResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r TestApiResponse) Failed() error {
	return r.E0
}

// TestApi implements Service. Primarily useful in a client.
func (e Endpoints) TestApi(ctx context.Context) (e0 error) {
	request := TestApiRequest{}
	response, err := e.TestApiEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(TestApiResponse).E0
}
