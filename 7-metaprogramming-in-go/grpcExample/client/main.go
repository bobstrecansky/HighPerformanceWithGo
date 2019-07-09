package main

import (
	"context"
	"log"
	"time"

	pb "github.com/HighPerformanceWithGo/7-metaprogramming-in-go/grpcExample/userinfo/userinfo"
	"google.golang.org/grpc"
)

const (
	defaultGrpcAddress = "localhost:50051"
	defaultUser        = "Gopher"
	defaultEmail       = "Gopher@example.com"
)

func main() {
	conn, err := grpc.Dial(defaultGrpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserInfoClient(conn)

	user := defaultUser
	email := defaultEmail
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.PrintUserInfo(ctx, &pb.UserInfoRequest{User: user, Email: email})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", r.Response)
}
