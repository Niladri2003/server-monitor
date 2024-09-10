package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string             `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName  string             `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Password  string             `bson:"password" json:"password" validate:"required,min=6"`
	UserRole  string             `bson:"user_role" json:"user_role" validate:"required,oneof=admin user"`
	Verified  bool               `bson:"verified" json:"verified"`
}
