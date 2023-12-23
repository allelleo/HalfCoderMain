package repository

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
