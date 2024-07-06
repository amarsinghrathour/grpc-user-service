package main

import (
	"google.golang.org/grpc"
	"grpc-user-service/factories"
	proto "grpc-user-service/proto/user_service"
	"grpc-user-service/services"
	"grpc-user-service/utilities"
	"log"
	"net"
)

func main() {

	// Parse command-line arguments to get the port number
	port := utilities.GetPortFromFlagsOrEnv()
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Create a factory to generate user repositories
	userRepositoryFactory := factories.UserFactory{}

	// Create a user repository instance using the factory
	userRepository := userRepositoryFactory.Create()

	// Create a new instance of UserServiceServer, passing in the user repository
	userServiceServer := services.NewUserServiceServer(userRepository)

	// Register the UserServiceServer with the gRPC server
	proto.RegisterUserServiceServer(grpcServer, userServiceServer)

	// Print a log message indicating the gRPC server is listening on port 50051
	log.Println("gRPC server listening on port 50051")

	// Start serving gRPC requests on the listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
