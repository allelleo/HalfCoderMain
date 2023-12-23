package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName         string          `gorm:"column:first_name"`
	LastName          string          `gorm:"column:last_name"`
	Email             string          `gorm:"column:email"`
	password          string          `gorm:"column:password"`
	UserName          string          `gorm:"column:user_name"`
	EmailValidationID int             `gorm:"column:email_validation_id"`
	EmailValidation   EmailValidation `gorm:"foreignkey:EmailValidationID"`
	Rating            RatingAll       `gorm:"foreignkey:RatingID"`
	RatingID          int             `gorm:"column:rating_id"`
	RatingMonth       RatingMonth     `gorm:"foreignkey:RatingMonthID"`
	RatingMonthID     int             `gorm:"column:rating_month_id"`
	RatingWeek        RatingWeek      `gorm:"foreignkey:RatingWeekID"`
	RatingWeekID      int             `gorm:"column:rating_week_id"`
	Profile           Profile         `gorm:"foreignkey:ProfileID"`
	ProfileID         int             `gorm:"column:profile_id"`
	Notifications     []Notification  `gorm:"many2many:user_notifications;"`
	Likes             []Like          `gorm:"many2many:user_likes;"`
	Subscribes        []Subscribe     `gorm:"many2many:user_subscribes;"`
	BookMarks         []BookMark      `gorm:"many2many:user_bookmarks;"`
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) CheckPassword(password string) bool {
	return u.password == password
}
