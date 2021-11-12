package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/joseavila-globant/bootcamp/userpb"
	"google.golang.org/grpc"
)

const (
	address   = "localhost:50051"
	defaultId = "123"
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
	id := defaultId
	if len(os.Args) > 1 {
		id = os.Args[1]
	}
	v, err := strconv.ParseInt(id, 10, 32)
	if err == nil {
		fmt.Printf("%q is not a valid id. \n", v)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &pb.UserRequest{Id: v})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("USER DATA: %s", r)
}
