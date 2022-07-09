// Copyright 2019 shimingyah. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// ee the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"mypoolserver/pb"
)

var port = flag.Int("port", 8888, "port number")

// server implements EchoServer.
type server struct{}

func (s *server) Say(context.Context, *pb.TestReq) (*pb.TestRsp, error) {
	fmt.Println("call  Say ... ")
	return &pb.TestRsp{Message: []byte("hello world")}, nil
}

func main() {
	flag.Parse()

	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTestsvrServer(s, &server{})
	fmt.Println("start server ...")

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


	fmt.Println("over server ...")
}
