package repository

type EmailValidation struct {
	ID   int    `gorm:"column:id"`
	Code string `gorm:"column:code"`
	Used bool   `gorm:"column:used"`
}
