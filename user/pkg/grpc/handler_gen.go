// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	endpoint "book/user/pkg/endpoint"
	pb "book/user/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	register        grpc.Handler
	userInfoById    grpc.Handler
	userInfoByPhone grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.UserServer {
	return &grpcServer{
		register:        makeRegisterHandler(endpoints, options["Register"]),
		userInfoById:    makeUserInfoByIdHandler(endpoints, options["UserInfoById"]),
		userInfoByPhone: makeUserInfoByPhoneHandler(endpoints, options["UserInfoByPhone"]),
	}
}
