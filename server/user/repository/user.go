package repository

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName         string          `gorm:"column:first_name"`
	LastName          string          `gorm:"column:last_name"`
	Email             string          `gorm:"column:email"`
	Password          string          `gorm:"column:password"`
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
	u.Password = password
	fmt.Println(u.Password)
}

func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}

func (u *User) JSON() fiber.Map {
	return fiber.Map{
		"id":         u.ID,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"email":      u.Email,
		"username":   u.UserName,
	}
}
