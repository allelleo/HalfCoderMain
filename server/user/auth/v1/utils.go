package v1

import (
	"errors"
	global "user/global"
	models "user/repository"

	"gorm.io/gorm"
)

func CheckUniqueUsername(username string) bool {
	db := global.GetDataBase()
	var user models.User
	res := db.Where("user_name = ?", username).First(&user)
	return errors.Is(res.Error, gorm.ErrRecordNotFound)
}
