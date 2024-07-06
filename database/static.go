package database

import (
	"grpc-user-service/models"
	"sync"
)

type Database struct {
	Users map[int32]*models.User
	sync.RWMutex
}

var instance *Database
var once sync.Once

// GetDatabase returns the singleton instance of the Database
func GetDatabase() *Database {
	once.Do(func() {
		instance = &Database{
			Users: map[int32]*models.User{
				1: {ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
				2: {ID: 2, FName: "John", City: "NY", Phone: 2345678901, Height: 6.0, Married: false},
				3: {ID: 3, FName: "Alice", City: "SF", Phone: 3456789012, Height: 5.5, Married: true},
				4: {ID: 4, FName: "Abram", City: "LA", Phone: 3456789065, Height: 5.2, Married: false},
			},
		}
	})
	return instance
}
