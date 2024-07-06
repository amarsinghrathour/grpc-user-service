package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"grpc-user-service/database"
	"grpc-user-service/models"
	"grpc-user-service/repositories"
	"testing"
)

// Implement other methods of Database interface as needed

func TestUserRepository_GetUserByID(t *testing.T) {
	db := database.GetDatabase() // Use the actual database singleton instance
	repo := &repositories.UserRepositoryImpl{
		Db: db,
	}

	testCases := []struct {
		name     string
		userID   int32
		expected *models.User
		err      error
	}{
		{
			name:     "Fetch existing user",
			userID:   1,
			expected: &models.User{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			err:      nil,
		},
		{
			name:     "Fetch non-existent user",
			userID:   100,
			expected: nil,
			err:      fmt.Errorf("user not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := repo.GetUserByID(tc.userID)
			if tc.err != nil {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tc.expected.FName, user.FName)
			}
		})
	}
}

func TestUserRepository_GetUsersByID(t *testing.T) {
	db := database.GetDatabase() // Use the actual database singleton instance
	repo := &repositories.UserRepositoryImpl{
		Db: db,
	}

	testCases := []struct {
		name     string
		userIDs  []int32
		expected []*models.User
		err      error
	}{
		{
			name:    "Fetch existing users",
			userIDs: []int32{1, 2},
			expected: []*models.User{
				{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
				{ID: 2, FName: "John", City: "NY", Phone: 2345678901, Height: 6.0, Married: false},
			},
			err: nil,
		},
		{
			name:     "Fetch users with non-existent IDs",
			userIDs:  []int32{100, 200},
			expected: nil,
			err:      fmt.Errorf("users not found for IDs: [100 200]"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			users, err := repo.GetUsersByID(tc.userIDs)
			if tc.err != nil {
				assert.Error(t, err)
				assert.Nil(t, users)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, users)
				assert.Equal(t, len(tc.expected), len(users))
			}
		})
	}
}

func TestUserRepository_SearchUsers(t *testing.T) {
	db := database.GetDatabase() // Use the actual database singleton instance
	repo := &repositories.UserRepositoryImpl{
		Db: db,
	}

	testCases := []struct {
		name     string
		city     string
		phone    int64
		married  bool
		expected []*models.User
		err      error
	}{
		{
			name:    "Search with valid criteria",
			city:    "LA",
			phone:   1234567890,
			married: true,
			expected: []*models.User{
				{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			},
			err: nil,
		},
		{
			name:     "Search with non-existent criteria",
			city:     "Chicago",
			phone:    9999999999,
			married:  false,
			expected: []*models.User{},
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			users, err := repo.SearchUsers(tc.city, tc.phone, tc.married)
			assert.NoError(t, err)
			assert.Equal(t, len(tc.expected), len(users))
		})
	}
}
