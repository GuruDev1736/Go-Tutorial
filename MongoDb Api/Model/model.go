package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	movie string             `json:"movie,omitempty"`
	watch bool               `json:"watched,omitempty"`
}
