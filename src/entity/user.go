package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username          string
	Password          string
	VideoList         []*Video
	FavoriteVideoList []*Video `gorm:"many2many:user_favorite_videos"`
}
