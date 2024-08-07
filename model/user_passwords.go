package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserPassword struct {
	ID             primitive.ObjectID
	UserID         primitive.ObjectID
	HashedPassword string
}
