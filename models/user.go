package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Roles    []string           `bson:"roles" json:"roles"`
	Status   uint8              `bson:"status" json:"status"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedAt int64 `bson:"created_at" json:"created_at"`
	UpdatedAt int64 `bson:"updated_at" json:"updated_at"`
}

type UserApi struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
}