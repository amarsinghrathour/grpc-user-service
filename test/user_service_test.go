package test

import (
	"context"
	proto "grpc-user-service/proto/user_service"
	"grpc-user-service/repositories"
	"grpc-user-service/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserServiceServer_GetUser(t *testing.T) {
	userRepository := repositories.NewUserRepository()
	service := services.NewUserServiceServer(userRepository)
	// Test scenarios
	var getUserTests = []struct {
		name       string
		id         int32
		expectedFn string
		expectErr  bool
	}{
		{"Existing User", 1, "Steve", false},
		{"Non-existent User", 100, "", true},
	}

	for _, tt := range getUserTests {
		t.Run(tt.name, func(t *testing.T) {
			req := &proto.GetUserRequest{Id: tt.id}
			resp, err := service.GetUser(context.Background(), req)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectedFn, resp.User.Fname)
			}
		})
	}
}

func TestUserServiceServer_GetUsers(t *testing.T) {
	userRepository := repositories.NewUserRepository()
	service := services.NewUserServiceServer(userRepository)
	// Test scenarios for GetUsers method
	var getUsersTests = []struct {
		name        string
		ids         []int32
		expectedLen int
		expectErr   bool
	}{
		{"Existing Users", []int32{1, 2}, 2, false},
		{"Non-existent Users", []int32{100, 200}, 0, true},
		{"Mixed Users", []int32{1, 100}, 0, true},
		{"Empty IDs", []int32{}, 0, false},
	}

	for _, tt := range getUsersTests {
		t.Run(tt.name, func(t *testing.T) {
			req := &proto.GetUsersRequest{Ids: tt.ids}
			resp, err := service.GetUsers(context.Background(), req)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Users, tt.expectedLen)
			}
		})
	}
}

func TestUserServiceServer_SearchUsers(t *testing.T) {
	userRepository := repositories.NewUserRepository()
	service := services.NewUserServiceServer(userRepository)
	// Test scenarios for SearchUsers method
	var searchUsersTests = []struct {
		name        string
		city        string
		phone       int64
		married     bool
		expectedLen int
	}{
		{"Valid Criteria", "LA", 1234567890, true, 1},
		{"Non-existent Criteria", "Chicago", 9999999999, false, 0},
		{"Only City", "LA", 0, false, 2},
		{"Only Phone", "", 1234567890, false, 1},
		{"Only Married", "", 0, true, 2},
		{"All Empty", "", 0, false, 4},
	}

	for _, tt := range searchUsersTests {
		t.Run(tt.name, func(t *testing.T) {
			req := &proto.SearchUsersRequest{
				City:    tt.city,
				Phone:   tt.phone,
				Married: tt.married,
			}
			resp, err := service.SearchUsers(context.Background(), req)

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Len(t, resp.Users, tt.expectedLen)
		})
	}
}
