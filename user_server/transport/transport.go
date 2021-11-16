package transport

import (
	"context"
	"log"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/joseavila-globant/bootcamp/user_server/endpoints"
	pb "github.com/joseavila-globant/bootcamp/userpb"
)

type gRPCServer struct {
	getUser gt.Handler
	pb.UnimplementedUserServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.UserServer {
	return &gRPCServer{
		getUser: gt.NewServer(
			endpoints.GetUser,
			decodeUserRequest,
			decodeUserDetails,
		),
		UnimplementedUserServer: pb.UnimplementedUserServer{},
	}
}

func (s *gRPCServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserDetails, error) {

	resp, err := s.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("could not get user GRPC : %v", err)
	}
	return resp, nil
}

func decodeUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserRequest)
	return endpoints.UserRequest{Id: req.Id}, nil
}

func decodeUserDetails(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.User)

	return &pb.UserDetails{Id: resp.Id, Name: resp.Name, Email: resp.Email, Pwd: &resp.Pwd, Age: resp.Age, Parents: nil}, nil
}
