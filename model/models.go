package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
	Salary  int                `json:"salary,omitempty"`
}

type Book struct {
	Id     int    `json:"id" gorm:"primary key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
	Pages  int    `json:"pages"`
}

type Handler struct {
	Path *gorm.DB
}
