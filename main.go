package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/kcwebapply/grpc-sample/protogen"
	"google.golang.org/grpc"
)

const (
	port = ":9999"
)

var (
	users = []*pb.User{}
)

type server struct{}

func (this server) GetUsers(ctx context.Context, req *empty.Empty) (*pb.Users, error) {
	//return &pb.Users{Users: []pb.User{&pb.User{Name: "kc", Address: "埼玉県"}}}, nil
	return &pb.Users{Users: users}, nil //&pb.Users{Users: []*pb.User{&pb.User{}}}, nil
}

func (this server) SaveUser(ctx context.Context, user *pb.User) (*empty.Empty, error) {
	users = append(users, user)
	//return &pb.Users{Users: []pb.User{&pb.User{Name: "kc", Address: "埼玉県"}}}, nil
	return new(empty.Empty), nil //&pb.Users{Users: []*pb.User{&pb.User{}}}, nil
}

//&pb.User{Name: "kc", Address: "埼玉県"}}

func main() {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, server{})
	if err := s.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
