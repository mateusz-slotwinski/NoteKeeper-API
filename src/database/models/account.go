package models

import primitive "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id       primitive.ObjectID `bson:"_id" form:"_id" json:"_id" validate:"required"`
	Name     string             `form:"name" json:"name,omitempty" validate:"required"`
	Email    string             `form:"email" json:"email,omitempty" validate:"required"`
	Login    string             `form:"login" json:"login,omitempty" validate:"required"`
	Password string             `form:"password" json:"password,omitempty" validate:"required"`
}
