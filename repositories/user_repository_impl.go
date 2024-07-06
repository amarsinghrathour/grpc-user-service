package repositories

import (
	"fmt"
	"grpc-user-service/database"
	"grpc-user-service/models"
)

type UserRepositoryImpl struct {
	Db *database.Database
}

// NewUserRepository returns a new instance of a UserRepository
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		Db: database.GetDatabase(),
	}
}

func (r *UserRepositoryImpl) GetUserByID(id int32) (*models.User, error) {
	r.Db.RLock()
	defer r.Db.RUnlock()

	// Check if the user exists in the database
	user, exists := r.Db.Users[id]
	if !exists {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetUsersByID(ids []int32) ([]*models.User, error) {
	r.Db.RLock()
	defer r.Db.RUnlock()

	var users []*models.User
	var notFoundIDs []int32

	for _, id := range ids {
		user, exists := r.Db.Users[id]
		if !exists {
			notFoundIDs = append(notFoundIDs, id)
		} else {
			users = append(users, user)
		}
	}

	if len(notFoundIDs) > 0 {
		return nil, fmt.Errorf("users not found for IDs: %v", notFoundIDs)
	}
	return users, nil
}

func (r *UserRepositoryImpl) SearchUsers(city string, phone int64, married bool) ([]*models.User, error) {
	r.Db.RLock()
	defer r.Db.RUnlock()

	var users []*models.User

	for _, user := range r.Db.Users {
		if (city == "" || user.City == city) &&
			(phone == 0 || user.Phone == phone) &&
			(!married || user.Married == married) {
			users = append(users, user)
		}
	}
	return users, nil
}
