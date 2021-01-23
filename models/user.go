package models

import "time"

type Favorite struct {
	ID         uint      `gorm:"primary_key"`
	UserID     string    `json:"user_id"`
	VideoID    string    `json:"video_id"`
	CreatedAt  time.Time `json:"-"`
	UpadatedAt time.Time `json:"-"`

	User User
}
