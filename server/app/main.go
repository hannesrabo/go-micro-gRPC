package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/hannesrabo/go-micro-gRPC/shared"
	"google.golang.org/grpc"
)

// This is the struct which we attach services to
type userService struct{}

func (s *userService) GetUserInfo(ctx context.Context, request *pb.UserInformationRequest) (*pb.UserInformation, error) {
	// Return the return message
	return &pb.UserInformation{Name: "Hannes Rabo"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("Failed to bind to port: %s\n", err)
	}

	server := grpc.NewServer()

	// Register our service			     (empty service interface)
	pb.RegisterUserServiceServer(server, &userService{})

	log.Printf("Server started!")

	// It there is any error while serving
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Error while serving: %v \n", err)
	}
}
