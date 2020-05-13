package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Info struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
}
