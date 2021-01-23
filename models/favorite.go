package models

import "time"

type User struct {
	ID         uint       `gorm:"primary_key"`
	UID        string     `json:"-"`
	CreatedAt  time.Time  `json:"-"`
	UpadatedAt time.Time  `json:"-"`
	DeletedAt  *time.Time `sql:"index"json:"-"`

	Favorites []Favorite
}
