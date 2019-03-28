package grpc

import (
	"book/curriculum/pkg/endpoint"
	"book/curriculum/pkg/grpc/pb"
	"context"
	"github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeCurriculumListHandler creates the handler logic
func makeCurriculumListHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CurriculumListEndpoint, decodeCurriculumListRequest, encodeCurriculumListResponse, options...)
}

// decodeCurriculumListResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeCurriculumListRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CurriculumListRequest)

	return endpoint.CurriculumListRequest{
		UserId:req.UserId,
	},nil
}

// encodeCurriculumListResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCurriculumListResponse(_ context.Context, r interface{}) (interface{}, error) {
	response := r.(endpoint.CurriculumListResponse)
	var curr = []*pb.CurriculumInfo{}
	for _,v := range response.C0{
		var exd = &pb.CurriculumInfo{
			Id:v.Id,
			UserId:v.UserId,
			Title:v.Title,
			Desc:v.Desc,
		}
		curr = append(curr,exd)
	}
	return &pb.CurriculumListReply{
		Info:curr,
	},nil

}
func (g *grpcServer) CurriculumList(ctx context1.Context, req *pb.CurriculumListRequest) (*pb.CurriculumListReply, error) {
	_, rep, err := g.curriculumList.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CurriculumListReply), nil
}

// makeCurriculumHandler creates the handler logic
func makeCurriculumHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CurriculumEndpoint, decodeCurriculumRequest, encodeCurriculumResponse, options...)
}

// decodeCurriculumResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeCurriculumRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CurriculumRequest)
	return endpoint.CurriculumRequest{
		Id:req.ID,
	},nil
}

// encodeCurriculumResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCurriculumResponse(_ context.Context, r interface{}) (interface{}, error) {
	response := r.(endpoint.CurriculumResponse)
	var info = &pb.CurriculumInfo{
		Id:response.C0.Id,
		UserId:response.C0.UserId,
		Title:response.C0.Title,
		Desc:response.C0.Desc,
	}
	return &pb.CurriculumReply{
		Info:info,
	},nil
}
func (g *grpcServer) Curriculum(ctx context1.Context, req *pb.CurriculumRequest) (*pb.CurriculumReply, error) {
	_, rep, err := g.curriculum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CurriculumReply), nil
}
