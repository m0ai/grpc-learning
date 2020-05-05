package main

import (
	"context"
	"log"
	"net"

	pb "github.com/m0ai/grpc-learn/protobuf"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server는 protobuf에서 정의된 함수의 인자로서 사용된다.
type server struct{}

// ProtoBuf의 IDL에 정의되어 있는 함수
// 함수의 인자와 리턴 값인 HelloRequest, HelloReply, 그리고 아래의 함수들은 모두
// protoc에서 생성된 skeleton 코드를 그대로 사용한다.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, Again!" + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
