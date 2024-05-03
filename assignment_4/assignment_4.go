package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) AddUser(ctx context.Context, in *pb.User) (*pb.AddUserResponse, error) {
	userID := int32(123)
	return &pb.AddUserResponse{UserId: userID}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	user := &pb.User{
		Id:    in.UserId,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	return user, nil
}

func (s *server) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users := []*pb.User{
		{Id: 1, Name: "John Doe", Email: "john@example.com"},
		{Id: 2, Name: "Jane Doe", Email: "jane@example.com"},
	}
	return &pb.ListUsersResponse{Users: users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	fmt.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
