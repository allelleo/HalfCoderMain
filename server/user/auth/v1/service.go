package v1

import (
	"fmt"
	"user/global"
	models "user/repository"

	"github.com/gofiber/fiber/v2"
)

func SignUpService(c *fiber.Ctx) error {
	data := new(SignUpModel)
	if err := c.BodyParser(data); err != nil {
		fmt.Println("BodyParser")
	}
	if data.IsValid() != 0 {
		fmt.Println("IsValid")
	}

	if !CheckUniqueUsername(data.UserName) {
		return UniqueUsernameError(c)
	}
	if !CheckEmailUnique(data.Email) {
		return UniqueEmailError(c)
	}

	var profile models.Profile = models.Profile{}
	var rating models.RatingAll = models.RatingAll{Score: 0}
	var ratingWeek models.RatingWeek = models.RatingWeek{Score: 0}
	var ratingMonth models.RatingMonth = models.RatingMonth{Score: 0}
	var emailValidation models.EmailValidation = models.EmailValidation{Code: "secret", Used: false}

	db := global.GetDataBase()
	db.Save(&profile)
	db.Save(&rating)
	db.Save(&ratingWeek)
	db.Save(&ratingMonth)
	db.Save(&emailValidation)

	var user models.User = models.User{
		FirstName:       data.FirstName,
		LastName:        data.LastName,
		Email:           data.Email,
		UserName:        data.UserName,
		Profile:         profile,
		Rating:          rating,
		RatingMonth:     ratingMonth,
		RatingWeek:      ratingWeek,
		EmailValidation: emailValidation,
	}
	user.SetPassword(data.Password)

	db.Save(&user)

	fmt.Println(user)

	return nil
}
