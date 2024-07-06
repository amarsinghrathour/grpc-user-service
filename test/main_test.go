package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"grpc-user-service/factories"
	"grpc-user-service/proto/user_service"
	"grpc-user-service/services"
	"net"
	"testing"
	"time"
)

const (
	address = "localhost:50051"
)

// startTestServer initializes and starts the gRPC server for testing.
func startTestServer(t *testing.T) *grpc.Server {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server object
	grpcServer := grpc.NewServer()

	// Create a factory to generate user repositories
	userRepositoryFactory := factories.UserFactory{}

	// Create a user repository instance using the factory
	userRepo := userRepositoryFactory.Create()

	userService := services.NewUserServiceServer(userRepo)

	// Attach the UserService service to the server
	user_service.RegisterUserServiceServer(grpcServer, userService)

	// Serve gRPC server
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()

	// Give the server a second to start
	time.Sleep(time.Second)

	return grpcServer
}

// TestUserService_GetUser tests the GetUser method of the UserService.
func TestUserService_GetUser(t *testing.T) {
	grpcServer := startTestServer(t)
	defer grpcServer.Stop()

	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*5))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := user_service.NewUserServiceClient(conn)

	// Define test cases
	testCases := []struct {
		id       int32
		expected string
		isError  bool
	}{
		{1, "Steve", false},
		{100, "", true},
	}

	for _, tc := range testCases {
		req := &user_service.GetUserRequest{Id: tc.id}
		resp, err := client.GetUser(context.Background(), req)
		if tc.isError {
			assert.Error(t, err)
			assert.Nil(t, resp)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tc.expected, resp.User.Fname)
		}
	}
}

// TestUserService_GetUsers tests the GetUsers method of the UserService.
func TestUserService_GetUsers(t *testing.T) {
	grpcServer := startTestServer(t)
	defer grpcServer.Stop()

	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*5))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := user_service.NewUserServiceClient(conn)

	// Define test cases
	testCases := []struct {
		ids      []int32
		expected int
		isError  bool
	}{
		{[]int32{1, 2}, 2, false},
		{[]int32{100, 200}, 0, true},
	}

	for _, tc := range testCases {
		req := &user_service.GetUsersRequest{Ids: tc.ids}
		resp, err := client.GetUsers(context.Background(), req)
		if tc.isError {
			assert.Error(t, err)
			assert.Nil(t, resp)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Len(t, resp.Users, tc.expected)
		}
	}
}

// TestUserService_SearchUsers tests the SearchUsers method of the UserService.
func TestUserService_SearchUsers(t *testing.T) {
	grpcServer := startTestServer(t)
	defer grpcServer.Stop()

	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*5))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := user_service.NewUserServiceClient(conn)

	// Define test cases
	testCases := []struct {
		city     string
		phone    int64
		married  bool
		expected int
	}{
		{"LA", 1234567890, true, 1},
		{"Chicago", 9999999999, false, 0},
	}

	for _, tc := range testCases {
		req := &user_service.SearchUsersRequest{City: tc.city, Phone: tc.phone, Married: tc.married}
		resp, err := client.SearchUsers(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Users, tc.expected)
	}
}
