package endpoint

import (
	service "book/user/pkg/service"
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
)

// RegisterRequest collects the request parameters for the Register method.
type RegisterRequest struct {
	Info service.UserInfo `json:"info"`
}

// RegisterResponse collects the response parameters for the Register method.
type RegisterResponse struct {
	I0 service.UserInfo `json:"i0"`
	E1 error `json:"e1"`
}

// MakeRegisterEndpoint returns an endpoint that invokes Register on the service.
func MakeRegisterEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		i0, e1 := s.Register(ctx, req.Info)
		return RegisterResponse{
			E1: e1,
			I0: i0,
		}, nil
	}
}

// Failed implements Failer.
func (r RegisterResponse) Failed() error {
	return r.E1
}

// UserInfoByIdRequest collects the request parameters for the UserInfoById method.
type UserInfoByIdRequest struct {
	Id int64 `json:"id"`
}

// UserInfoByIdResponse collects the response parameters for the UserInfoById method.
type UserInfoByIdResponse struct {
	U0 service.UserInfo `json:"u0"`
	E1 error            `json:"e1"`
}

// MakeUserInfoByIdEndpoint returns an endpoint that invokes UserInfoById on the service.
func MakeUserInfoByIdEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserInfoByIdRequest)
		u0, e1 := s.UserInfoById(ctx, req.Id)
		return UserInfoByIdResponse{
			E1: e1,
			U0: u0,
		}, nil
	}
}

// Failed implements Failer.
func (r UserInfoByIdResponse) Failed() error {
	return r.E1
}

// UserInfoByPhoneRequest collects the request parameters for the UserInfoByPhone method.
type UserInfoByPhoneRequest struct {
	Phone int64 `json:"phone"`
}

// UserInfoByPhoneResponse collects the response parameters for the UserInfoByPhone method.
type UserInfoByPhoneResponse struct {
	U0 service.UserInfo `json:"u0"`
	E1 error            `json:"e1"`
}

// MakeUserInfoByPhoneEndpoint returns an endpoint that invokes UserInfoByPhone on the service.
func MakeUserInfoByPhoneEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserInfoByPhoneRequest)
		u0, e1 := s.UserInfoByPhone(ctx, req.Phone)
		return UserInfoByPhoneResponse{
			E1: e1,
			U0: u0,
		}, nil
	}
}

// Failed implements Failer.
func (r UserInfoByPhoneResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Register implements Service. Primarily useful in a client.
func (e Endpoints) Register(ctx context.Context, info service.UserInfo) (i0 service.UserInfo, e1 error) {
	request := RegisterRequest{Info: info}
	response, err := e.RegisterEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RegisterResponse).I0, response.(RegisterResponse).E1
}

// UserInfoById implements Service. Primarily useful in a client.
func (e Endpoints) UserInfoById(ctx context.Context, id int64) (u0 service.UserInfo, e1 error) {
	request := UserInfoByIdRequest{Id: id}
	response, err := e.UserInfoByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UserInfoByIdResponse).U0, response.(UserInfoByIdResponse).E1
}

// UserInfoByPhone implements Service. Primarily useful in a client.
func (e Endpoints) UserInfoByPhone(ctx context.Context, phone int64) (u0 service.UserInfo, e1 error) {
	request := UserInfoByPhoneRequest{Phone: phone}
	response, err := e.UserInfoByPhoneEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UserInfoByPhoneResponse).U0, response.(UserInfoByPhoneResponse).E1
}
