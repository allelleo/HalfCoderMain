package global

import (
	user "user/repository"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDataBase() {
	db, err := gorm.Open(sqlite.Open(DATABASE_URL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&user.BookMarkType{},
		&user.BookMark{},

		&user.EmailValidation{},

		&user.LikeType{},
		&user.Like{},

		&user.SubscribeType{},
		&user.Subscribe{},

		&user.NotificationType{},
		&user.Notification{},

		&user.Profile{},

		&user.RatingAll{},
		&user.RatingMonth{},
		&user.RatingWeek{},

		&user.User{},
	)
}

func GetDataBase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DATABASE_URL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
