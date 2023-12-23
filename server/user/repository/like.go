package repository

import "gorm.io/gorm"

type LikeType struct {
	ID    int    `gorm:"column:id"`
	Title string `gorm:"column:title"`
}

type Like struct {
	gorm.Model
	LikeType   LikeType `gorm:"foreignkey:LikeTypeID"`
	LikeTypeID int      `gorm:"column:like_type_id"`
	LikeTo     int      `gorm:"column:like_to"`
	UserID     int      `gorm:"column:user_id"`
}
