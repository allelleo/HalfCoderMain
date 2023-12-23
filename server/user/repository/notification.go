package repository

import "gorm.io/gorm"

type NotificationType struct {
	ID    int    `gorm:"column:id"`
	Title string `gorm:"column:title"`
}

type Notification struct {
	gorm.Model
	Title              string           `gorm:"column:title"`
	Content            string           `gorm:"column:content"`
	NotificationType   NotificationType `gorm:"foreignkey:NotificationTypeID"`
	NotificationTypeID int              `gorm:"column:notification_type_id"`
	UserID             int              `gorm:"column:user_id"`
}
