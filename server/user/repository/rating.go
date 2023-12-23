package repository

import "github.com/gofiber/fiber/v2"

type RatingAll struct {
	ID    int `gorm:"column:id"`
	Score int `gorm:"column:score"`
}

type RatingMonth struct {
	ID    int `gorm:"column:id"`
	Score int `gorm:"column:score"`
}

type RatingWeek struct {
	ID    int `gorm:"column:id"`
	Score int `gorm:"column:score"`
}

func (r *RatingAll) JSON() fiber.Map {
	return fiber.Map{
		"id":    r.ID,
		"score": r.Score,
	}
}
func (r *RatingMonth) JSON() fiber.Map {
	return fiber.Map{
		"id":    r.ID,
		"score": r.Score,
	}
}
func (r *RatingWeek) JSON() fiber.Map {
	return fiber.Map{
		"id":    r.ID,
		"score": r.Score,
	}
}
