package main

import (
	"context"
	"log"
	"net"

	pb "github.com/HighPerformanceWithGo/7-metaprogramming-in-go/grpcExample/userinfo/userinfo"
	"google.golang.org/grpc"
)

type userInfoServer struct{}

func (s *userInfoServer) PrintUserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	log.Printf("%s %s", in.User, in.Email)
	return &pb.UserInfoResponse{Response: "User Info: User Name: " + in.User + " User Email: " + in.Email}, nil
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserInfoServer(s, &userInfoServer{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Couldn't create Server: %v", err)
	}
}
