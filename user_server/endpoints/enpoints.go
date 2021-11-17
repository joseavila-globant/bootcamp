package endpoints

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/joseavila-globant/bootcamp/user_server/service"
	pb "github.com/joseavila-globant/bootcamp/userpb"
)

type Endpoints struct {
	GetUser endpoint.Endpoint
}

type UserRequest struct {
	Id int64 `json:"id"`
}

var logs log.Logger

type User struct {
	Id      int64  `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Pwd     string `json:"pwd,omitempty"`
	Parents []User `json:"parents,omitempty"`
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GetUser: makeGetUserEndpoint(s),
	}
}

func makeGetUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserRequest)
		logs.Log("Msg", "trying to get user endpoints")

		User, err := s.GetUser(ctx, &pb.UserRequest{Id: req.Id})

		if err != nil {

			logs.Log("error:", err)

		}
		return User, nil

	}
}
