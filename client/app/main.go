package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/hannesrabo/go-micro-gRPC/shared"
	"google.golang.org/grpc"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellu !! You know this is amazing!")
}

const (
	address = "grpc-server:8081"
)

func main() {

	log.Printf("Dialing: %v\n", address)
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	defer connection.Close()

	if err != nil {
		log.Fatalf("Faild to open connection to grpc server: \n%v\n", err)
	}

	// Create the client
	server := pb.NewUserServiceClient(connection)

	// Create a timeout object
	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := server.GetUserInfo(context, &pb.UserInformationRequest{Id: 123})

	if err != nil {
		log.Fatalf("Failed to perform request: \n%v\n", err)
	}

	log.Printf("Response: %v\n", response)

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
