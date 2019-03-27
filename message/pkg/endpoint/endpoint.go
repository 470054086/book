package endpoint

import (
	service "book/message/pkg/service"
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
)

// SendEmailMessageRequest collects the request parameters for the SendEmailMessage method.
type SendEmailMessageRequest struct {
	Email   string `json:"email"`
	Text    string `json:"text"`
	Content string `json:"content"`
}

// SendEmailMessageResponse collects the response parameters for the SendEmailMessage method.
type SendEmailMessageResponse struct {
	Err error `json:"err"`
}

// MakeSendEmailMessageEndpoint returns an endpoint that invokes SendEmailMessage on the service.
func MakeSendEmailMessageEndpoint(s service.MessageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendEmailMessageRequest)
		err := s.SendEmailMessage(ctx, req.Email, req.Text, req.Content)
		return SendEmailMessageResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r SendEmailMessageResponse) Failed() error {
	return r.Err
}

// SendPhoneMessageRequest collects the request parameters for the SendPhoneMessage method.
type SendPhoneMessageRequest struct {
	Phone   string `json:"phone"`
	Content string `json:"content"`
}

// SendPhoneMessageResponse collects the response parameters for the SendPhoneMessage method.
type SendPhoneMessageResponse struct {
	Err error `json:"err"`
}

// MakeSendPhoneMessageEndpoint returns an endpoint that invokes SendPhoneMessage on the service.
func MakeSendPhoneMessageEndpoint(s service.MessageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendPhoneMessageRequest)
		err := s.SendPhoneMessage(ctx, req.Phone, req.Content)
		return SendPhoneMessageResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r SendPhoneMessageResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SendEmailMessage implements Service. Primarily useful in a client.
func (e Endpoints) SendEmailMessage(ctx context.Context, email string, text string, content string) (err error) {
	request := SendEmailMessageRequest{
		Content: content,
		Email:   email,
		Text:    text,
	}
	response, err := e.SendEmailMessageEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendEmailMessageResponse).Err
}

// SendPhoneMessage implements Service. Primarily useful in a client.
func (e Endpoints) SendPhoneMessage(ctx context.Context, phone string, content string) (err error) {
	request := SendPhoneMessageRequest{
		Content: content,
		Phone:   phone,
	}
	response, err := e.SendPhoneMessageEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendPhoneMessageResponse).Err
}
