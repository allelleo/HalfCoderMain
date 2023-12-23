package repository

import "gorm.io/gorm"

type BookMarkType struct {
	ID    int    `gorm:"column:id"`
	Title string `gorm:"column:title"`
}

type BookMark struct {
	gorm.Model
	BookMarkTypeID int          `gorm:"column:bookmark_type_id"`
	BookMarkType   BookMarkType `gorm:"foreignkey:BookMarkTypeID"`
	BookMarkTo     int          `gorm:"column:bookmark_to"`
	UserID         int          `gorm:"column:user_id"`
}
