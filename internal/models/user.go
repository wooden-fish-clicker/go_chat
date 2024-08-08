package models

import "github.com/wooden-fish-clicker/chat/pkg/db"

type User struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Country  string `bson:"country"`
	Points   int    `bson:"points"`

	db.BaseModel `bson:",inline"`
}
