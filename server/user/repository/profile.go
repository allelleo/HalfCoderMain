package repository

import "time"

type Profile struct {
	ID        int       `gorm:"column:id"`
	Work      string    `gorm:"column:work"`
	Sex       string    `gorm:"column:sex"`
	Birtday   time.Time `gorm:"column:birthday"`
	Education string    `gorm:"column:education"`
	Hobby     string    `gorm:"column:hobby"`
	Quote     string    `gorm:"column:quote"`
	Phone     string    `gorm:"column:phone"`
	Country   string    `gorm:"column:country"`
	WebSite   string    `gorm:"column:website"`
}
