package grpc

import (
	endpoint "book/user/pkg/endpoint"
	pb "book/user/pkg/grpc/pb"
	"context"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeRegisterHandler creates the handler logic
func makeRegisterHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, options...)
}

// decodeRegisterResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeRegisterRequest(_ context.Context, r interface{}) (interface{}, error) {
	//return nil, errors.New("'User' Decoder is not impelemented")
	req := r.(*pb.RegisterRequest)
	return endpoint.RegisterRequest{
		Info:req.User,
	},nil
}

// encodeRegisterResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeRegisterResponse(_ context.Context, r interface{}) (interface{}, error) {
	response := r.(endpoint.RegisterResponse)
	return &pb.RegisterReply{
		User: response.I0,
	},response.E1
}
func (g *grpcServer) Register(ctx context1.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	_, rep, err := g.register.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RegisterReply), nil
}

// makeUserInfoByIdHandler creates the handler logic
func makeUserInfoByIdHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UserInfoByIdEndpoint, decodeUserInfoByIdRequest, encodeUserInfoByIdResponse, options...)
}

// decodeUserInfoByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeUserInfoByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.UserInfoByIdRequest)
	return endpoint.UserInfoByIdRequest{
		Id:req.Id,
	},nil
}

// encodeUserInfoByIdResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUserInfoByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	//return nil, errors.New("'User' Encoder is not impelemented")
	response := r.(endpoint.UserInfoByIdResponse)
	return &pb.UserInfoByIdReply{
		User:response.U0,
	},response.E1


}
func (g *grpcServer) UserInfoById(ctx context1.Context, req *pb.UserInfoByIdRequest) (*pb.UserInfoByIdReply, error) {
	_, rep, err := g.userInfoById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UserInfoByIdReply), nil
}

// makeUserInfoByPhoneHandler creates the handler logic
func makeUserInfoByPhoneHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UserInfoByPhoneEndpoint, decodeUserInfoByPhoneRequest, encodeUserInfoByPhoneResponse, options...)
}

// decodeUserInfoByPhoneResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeUserInfoByPhoneRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.UserInfoByPhoneRequest)
	return endpoint.UserInfoByPhoneRequest{
		Phone:req.Phone,
	},nil
}

// encodeUserInfoByPhoneResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUserInfoByPhoneResponse(_ context.Context, r interface{}) (interface{}, error) {
	response := r.(endpoint.UserInfoByPhoneResponse)
	return &pb.UserInfoByPhoneReply{
		User:response.U0,
	},response.E1
}
func (g *grpcServer) UserInfoByPhone(ctx context1.Context, req *pb.UserInfoByPhoneRequest) (*pb.UserInfoByPhoneReply, error) {
	_, rep, err := g.userInfoByPhone.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UserInfoByPhoneReply), nil
}
