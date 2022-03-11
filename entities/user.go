package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `json:"name" form:"name"`
	Email    string             `json:"email" form:"email"`
	Password string             `json:"password" form:"password"`
}
