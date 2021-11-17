package main

import (
	"context"
	"log"
	"time"

	pb "github.com/joseavila-globant/bootcamp/userpb"
	"google.golang.org/grpc"
)

const (
	address   string = "localhost:50051"
	defaultId string = "123"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	// Contact the server and print out its response.
	// id := defaultId
	// if len(os.Args) > 1 {
	// 	id = os.Args[1]
	// }
	// v, err := strconv.ParseInt(id, 10, 64)
	// if err == nil {
	// 	fmt.Printf("Not a valid id. \n erros = %v", err)
	// }
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &pb.UserRequest{Id: 123})
	if err != nil {
		log.Fatalf("could not Get User: %v", err)
	}
	log.Printf("USER DATA: %s", r)
}
