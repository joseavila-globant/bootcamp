/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/joseavila-globant/bootcamp/user_server/endpoints"
	"github.com/joseavila-globant/bootcamp/user_server/service"
	transport "github.com/joseavila-globant/bootcamp/user_server/transport"
	pb "github.com/joseavila-globant/bootcamp/userpb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement serServer.
// type server struct {
// 	pb.UnimplementedUserServer
// }

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	getUserService := service.NewService(logger)
	getUserEndpoint := endpoints.MakeEndpoints(getUserService)
	grpcServer := transport.NewGRPCServer(getUserEndpoint, logger)
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Log("error listening in port ", port, "error %v", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterUserServer(baseServer, grpcServer)

		level.Info(logger).Log("msg", "Server started")
		baseServer.Serve(grpcListener)

	}()
	level.Error(logger).Log("exit", <-errs)

}
