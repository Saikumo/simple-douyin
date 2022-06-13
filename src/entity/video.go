package entity

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	UserId   uint
	PlayUrl  string
	CoverUrl string
	title    string
}
