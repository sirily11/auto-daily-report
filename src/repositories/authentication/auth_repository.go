package auth

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepositoryInterface interface {
	// Create a new user

}

// AuthRepository is the implementation of the AuthRepositoryInterface
type AuthRepository struct {
	userCollection *mongo.Collection
}
