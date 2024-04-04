package model

import "time"

type User struct {
	UserID    int
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	MetaInfo  string
}
