package user

import (
	"go-mongodb/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User interface {
	Get() ([]entities.User, error)
	GetByEmail(string) (entities.User, error)
	Create(entities.User) (entities.User, error)
	Update(entities.User, primitive.ObjectID) error
	Delete(primitive.ObjectID) error
}
