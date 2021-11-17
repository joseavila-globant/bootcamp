package service

import (
	"context"

	"github.com/go-kit/kit/log"
	pb "github.com/joseavila-globant/bootcamp/userpb"
)

type service struct {
	logger log.Logger
}

type Service interface {
	GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserDetails, error)
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}

}

func (s service) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserDetails, error) {
	s.logger.Log("Received: %v", in.GetId())
	s.logger.Log("Msg", "Service id")
	Pass := "123456"
	return &pb.UserDetails{
		Id:      in.GetId(),
		Name:    "Jose avila",
		Email:   "jose.avila@globant.com",
		Pwd:     &Pass,
		Age:     29,
		Parents: nil,
	}, nil
}
