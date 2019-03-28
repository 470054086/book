package grpc

import (
	"book/user/pkg/endpoint"
	"book/user/pkg/grpc/pb"
	"context"
	"github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	"book/user/pkg/service"
)

// makeRegisterHandler creates the handler logic
func makeRegisterHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, options...)
}

// decodeRegisterResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeRegisterRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.RegisterRequest)
	//进行数据的转换
	var user = service.UserInfo{
		Id:req.User.Id,
		Phone:req.User.Phone,
		Password:req.User.Password,
		Age:req.User.Age,

	}
	return endpoint.RegisterRequest{
		User:user,
	},nil
}

// encodeRegisterResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeRegisterResponse(_ context.Context, r interface{}) (interface{}, error) {
	response := r.(endpoint.RegisterResponse)
	if response.U0 == nil {
		return nil,response.E1
	}
	//进行数据的转换
	var user = &pb.UserInfo{
		Id:response.U0.Id,
		Phone:response.U0.Phone,
		Password:response.U0.Password,
		Age:response.U0.Age,
	}
	return &pb.RegisterReply{
		User:user,
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
	response := r.(endpoint.UserInfoByIdResponse)
	var user = &pb.UserInfo{
		Id:response.U0.Id,
		Phone:response.U0.Phone,
		Password:response.U0.Password,
		Age:response.U0.Age,
	}
	return &pb.UserInfoByIdReply{
		User:user,
	},nil
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
	var user = &pb.UserInfo{
		Id:response.U0.Id,
		Phone:response.U0.Phone,
		Password:response.U0.Password,
		Age:response.U0.Age,
	}
	return &pb.UserInfoByPhoneReply{
		User:user,
	},nil
}
func (g *grpcServer) UserInfoByPhone(ctx context1.Context, req *pb.UserInfoByPhoneRequest) (*pb.UserInfoByPhoneReply, error) {
	_, rep, err := g.userInfoByPhone.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UserInfoByPhoneReply), nil
}
