package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Todo struct {
	Id        bson.ObjectID `bson:"_id,omitempty"`
	Title     string        `bson:"title,omitempty"`
	Content   string        `bson:"content,omitempty"`
	CreatedAt int64         `bson:"created_at,omitempty"`
}
