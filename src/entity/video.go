package entity

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	UserId           uint
	PlayUrl          string
	CoverUrl         string
	Title            string
	FavoriteUserList []*User `gorm:"many2many:user_favorite_videos"`
	FavoriteCount    uint
}
