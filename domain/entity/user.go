package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Age      int8               `bson:"age"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}
