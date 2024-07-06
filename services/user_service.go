package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	proto "grpc-user-service/proto/user_service"
	"grpc-user-service/repositories"
	"log"
	"os"
)

// UserServiceServer is the server that provides user services.
type UserServiceServer struct {
	proto.UnimplementedUserServiceServer                             // Embeds the unimplemented gRPC service interface
	repo                                 repositories.UserRepository // Repository for user data operations
	logger                               *log.Logger                 // Logger for logging service operations
}

// NewUserServiceServer returns a new UserServiceServer instance.
func NewUserServiceServer(repo repositories.UserRepository) *UserServiceServer {
	return &UserServiceServer{
		repo:   repo,
		logger: log.New(os.Stdout, "userService: ", log.LstdFlags),
	}
}

// GetUser fetches user details by ID.
func (s *UserServiceServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	s.logger.Printf("GetUser request received for ID: %d\n", req.Id)

	user, err := s.repo.GetUserByID(req.Id)
	if err != nil {
		s.logger.Printf("Failed to fetch user for ID %d: %v\n", req.Id, err)
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	s.logger.Printf("User fetched successfully for ID: %d\n", req.Id)

	return &proto.GetUserResponse{
		User: &proto.User{
			Id:      user.ID,
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		},
	}, nil
}

// GetUsers fetches user details by a list of IDs.
func (s *UserServiceServer) GetUsers(ctx context.Context, req *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	s.logger.Printf("GetUsers request received for IDs: %v\n", req.Ids)

	users, err := s.repo.GetUsersByID(req.Ids)
	if err != nil {
		s.logger.Printf("Failed to fetch users for IDs %v: %v\n", req.Ids, err)
		return nil, status.Errorf(codes.Internal, "Failed to fetch users")
	}

	s.logger.Printf("Users fetched successfully for IDs: %v\n", req.Ids)

	var protoUsers []*proto.User
	for _, user := range users {
		protoUsers = append(protoUsers, &proto.User{
			Id:      user.ID,
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}

	return &proto.GetUsersResponse{Users: protoUsers}, nil
}

// SearchUsers searches users by criteria.
func (s *UserServiceServer) SearchUsers(ctx context.Context, req *proto.SearchUsersRequest) (*proto.SearchUsersResponse, error) {
	s.logger.Printf("SearchUsers request received with criteria: City=%s, Phone=%v, Married=%v\n", req.City, req.Phone, req.Married)

	users, err := s.repo.SearchUsers(req.City, req.Phone, req.Married)
	if err != nil {
		s.logger.Printf("Failed to search users with criteria: City=%s, Phone=%v, Married=%v: %v\n", req.City, req.Phone, req.Married, err)
		return nil, status.Errorf(codes.Internal, "Failed to search users")
	}

	s.logger.Printf("Users searched successfully with criteria: City=%s, Phone=%v, Married=%v\n", req.City, req.Phone, req.Married)

	var protoUsers []*proto.User
	for _, user := range users {
		protoUsers = append(protoUsers, &proto.User{
			Id:      user.ID,
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}

	return &proto.SearchUsersResponse{Users: protoUsers}, nil
}
