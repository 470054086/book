package grpc

import (
	"book/message/pkg/endpoint"
	"book/message/pkg/grpc/pb"
	"context"
	"errors"
	"github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeSendEmailMessageHandler creates the handler logic
func makeSendEmailMessageHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendEmailMessageEndpoint, decodeSendEmailMessageRequest, encodeSendEmailMessageResponse, options...)
}

// decodeSendEmailMessageResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeSendEmailMessageRequest(_ context.Context, r interface{}) (interface{}, error) {
	//return nil, errors.New("'Message' Decoder is not impelemented")
	req := r.(*pb.SendEmailMessageRequest)
	return endpoint.SendEmailMessageRequest{
		Email:req.Email,
		Text:req.Text,
		Content:req.Content,
	},nil
}

// encodeSendEmailMessageResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendEmailMessageResponse(_ context.Context, r interface{}) (interface{}, error) {
	//return nil, errors.New("'Message' Encoder is not impelemented")
	response := r.(endpoint.SendEmailMessageResponse)
	return &pb.SendEmailMessageReply{},response.Err
}
func (g *grpcServer) SendEmailMessage(ctx context1.Context, req *pb.SendEmailMessageRequest) (*pb.SendEmailMessageReply, error) {
	_, rep, err := g.sendEmailMessage.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendEmailMessageReply), nil
}

// makeSendPhoneMessageHandler creates the handler logic
func makeSendPhoneMessageHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendPhoneMessageEndpoint, decodeSendPhoneMessageRequest, encodeSendPhoneMessageResponse, options...)
}

// decodeSendPhoneMessageResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeSendPhoneMessageRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Message' Decoder is not impelemented")
}

// encodeSendPhoneMessageResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendPhoneMessageResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Message' Encoder is not impelemented")
}
func (g *grpcServer) SendPhoneMessage(ctx context1.Context, req *pb.SendPhoneMessageRequest) (*pb.SendPhoneMessageReply, error) {
	_, rep, err := g.sendPhoneMessage.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendPhoneMessageReply), nil
}
