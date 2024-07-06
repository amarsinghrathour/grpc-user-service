package repositories

import "grpc-user-service/models"

// UserRepository defines the methods that any
// data storage provider needs to implement to get
// and store users
type UserRepository interface {
	GetUserByID(id int32) (*models.User, error)
	GetUsersByID(ids []int32) ([]*models.User, error)
	SearchUsers(city string, phone int64, married bool) ([]*models.User, error)
}
