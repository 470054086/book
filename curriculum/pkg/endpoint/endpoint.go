package endpoint

import (
	service "book/curriculum/pkg/service"
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CurriculumListRequest collects the request parameters for the CurriculumList method.
type CurriculumListRequest struct {
	UserId int64 `json:"user_id"`
}

// CurriculumListResponse collects the response parameters for the CurriculumList method.
type CurriculumListResponse struct {
	C0 []service.CurriculumInfo `json:"c0"`
	E1 error            `json:"e1"`
}

// MakeCurriculumListEndpoint returns an endpoint that invokes CurriculumList on the service.
func MakeCurriculumListEndpoint(s service.CurriculumService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CurriculumListRequest)
		c0, e1 := s.CurriculumList(ctx, req.UserId)
		return CurriculumListResponse{
			C0: c0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r CurriculumListResponse) Failed() error {
	return r.E1
}

// CurriculumRequest collects the request parameters for the Curriculum method.
type CurriculumRequest struct {
	Id int64 `json:"id"`
}

// CurriculumResponse collects the response parameters for the Curriculum method.
type CurriculumResponse struct {
	C0 *service.CurriculumInfo `json:"c0"`
	E1 error           `json:"e1"`
}

// MakeCurriculumEndpoint returns an endpoint that invokes Curriculum on the service.
func MakeCurriculumEndpoint(s service.CurriculumService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CurriculumRequest)
		c0, e1 := s.Curriculum(ctx, req.Id)
		return CurriculumResponse{
			C0: c0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r CurriculumResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CurriculumList implements Service. Primarily useful in a client.
func (e Endpoints) CurriculumList(ctx context.Context, userId int64) (c0 []service.CurriculumInfo, e1 error) {
	request := CurriculumListRequest{UserId: userId}
	response, err := e.CurriculumListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CurriculumListResponse).C0, response.(CurriculumListResponse).E1
}

// Curriculum implements Service. Primarily useful in a client.
func (e Endpoints) Curriculum(ctx context.Context, id int64) (c0 *service.CurriculumInfo, e1 error) {
	request := CurriculumRequest{Id: id}
	response, err := e.CurriculumEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CurriculumResponse).C0, response.(CurriculumResponse).E1
}
