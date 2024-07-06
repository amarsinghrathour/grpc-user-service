package factories

import "grpc-user-service/repositories"

// UserFactory is responsible for creating new instances of UserRepository
type UserFactory struct{}

// Create returns a new instance of a UserRepository
func (f *UserFactory) Create() repositories.UserRepository {
	return repositories.NewUserRepository()
}
