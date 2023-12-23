package v1

import (
	"fmt"
	"user/global"
	models "user/repository"

	"github.com/gofiber/fiber/v2"
)

func MeService(c *fiber.Ctx) error {
	data := new(MeModel)
	if err := c.BodyParser(data); err != nil {
		return JSONParsingError(c)
	}
	if data.IsValid() != 0 {
	}

	payload, _ := global.DecodeToken(data.Token)
	fmt.Println(payload)

	var user_id int = int(payload["user_id"].(float64))
	fmt.Println(user_id)
	var user models.User

	db := global.GetDataBase() // Получение экземпляра базы данных

	db.Preload("Profile").Preload("Rating").Preload("RatingMonth").Preload("RatingWeek").Preload("Notifications").First(&user, user_id)

	fmt.Println(user.Notifications)
	for i := 0; i < len(user.Notifications); i++ {
		var n models.NotificationType
		db.First(&n, user.Notifications[i].NotificationTypeID)
		fmt.Println(user.Notifications[i].Title, n.Title)
	}

	return c.Status(200).JSON(fiber.Map{
		"user":    user.JSON(),
		"profile": user.Profile.JSON(),
		"ratings": fiber.Map{
			"all":   user.Rating.JSON(),
			"week":  user.RatingWeek.JSON(),
			"month": user.RatingMonth.JSON(),
		},
	})

}
