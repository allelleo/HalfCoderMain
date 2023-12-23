package repository

import "user/repository/utils"

type EmailValidation struct {
	ID   int    `gorm:"column:id"`
	Code string `gorm:"column:code"`
	Used bool   `gorm:"column:used"`
}

func GenerateEmailValidation() EmailValidation {
	return EmailValidation{
		Code: utils.GenerateCode(30),
		Used: false,
	}
}
