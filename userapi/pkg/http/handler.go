package http

import (
	"book/userapi/pkg/endpoint"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	http1 "github.com/go-kit/kit/transport/http"
)

// makeRegisterHandler creates the handler logic
func makeRegisterHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/", http1.NewServer(endpoints.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, options...))
}

// decodeRegisterRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRegisterResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRegisterResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeLoginHandler creates the handler logic
func makeLoginHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/login", http1.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...))
}

// decodeLoginRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeLoginResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {

	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserInfoHandler creates the handler logic
func makeGetUserInfoHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-user-info", http1.NewServer(endpoints.GetUserInfoEndpoint, decodeGetUserInfoRequest, encodeGetUserInfoResponse, options...))
}

// decodeGetUserInfoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUserInfoRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUserInfoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserInfoResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(err2code(err))
	data := make(map[string]interface{})
	json.NewEncoder(w).Encode(Wrapper{Error: err.Error(), Data: data, Message: ""})
}
func ErrorDecoder(r *http.Response) error {
	var w Wrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type Wrapper struct {
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// makeTestHandler creates the handler logic
func makeTestHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/test", http1.NewServer(endpoints.TestEndpoint, decodeTestRequest, encodeTestResponse, options...))
}

// decodeTestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeTestRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.TestRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeTestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeTestResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeTestApiHandler creates the handler logic
func makeTestApiHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/test-api", http1.NewServer(endpoints.TestApiEndpoint, decodeTestApiRequest, encodeTestApiResponse, options...))
}

// decodeTestApiRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeTestApiRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.TestApiRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeTestApiResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeTestApiResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
