package models

import primitive "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	Id      primitive.ObjectID `bson:"_id" form:"_id" json:"_id" validate:"required"`
	Title   string             `form:"title" json:"title,omitempty"`
	Content string             `form:"content" json:"content,omitempty" validate:"required"`
	Author  string             `form:"author" json:"author,omitempty" validate:"required"`
}
