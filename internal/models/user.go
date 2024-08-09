package models

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Email     string    `bson:"email"`
	Account   string    `bson:"account"`
	Password  string    `bson:"password"`
	UserInfo  UserInfo  `bson:"user_info"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

type UserInfo struct {
	Name    string `bson:"name"`
	Country string `bson:"country"`
	Points  int    `bson:"points"`
	Hp      int    `bson:"hp"`
}
