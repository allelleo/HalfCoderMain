package repository

import "gorm.io/gorm"

type SubscribeType struct {
	ID    int    `gorm:"column:id"`
	Title string `gorm:"column:title"`
}

type Subscribe struct {
	gorm.Model
	SubscribeType   SubscribeType `gorm:"foreignkey:SubscribeTypeID"`
	SubscribeTypeID int           `gorm:"column:subscribe_type_id"`
	SubscribeTo     int           `gorm:"column:subscribe_to"`
	UserID          int           `gorm:"column:user_id"`
}
