package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserPassword struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	UserID         primitive.ObjectID
	HashedPassword string
}
