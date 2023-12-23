package repository

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

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

func CreateProfile() Profile {
	return Profile{
		Work:      "Не указанно",
		Sex:       "Не указанно",
		Education: "Не указанно",
		Hobby:     "Не указанно",
		Quote:     "Привет! я пользуюсь сайтом HalfCoder",
		Phone:     "Не указанно",
		Country:   "Не указанно",
		WebSite:   "Не указанно",
	}
}

func (p *Profile) JSON() fiber.Map {
	return fiber.Map{
		"id":        p.ID,
		"work":      p.Work,
		"sex":       p.Sex,
		"birthday":  p.Birtday,
		"education": p.Education,
		"hobby":     p.Hobby,
		"quote":     p.Quote,
		"phone":     p.Phone,
		"country":   p.Country,
		"website":   p.WebSite,
	}
}
