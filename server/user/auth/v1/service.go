package v1

import (
	"user/global"
	models "user/repository"

	"github.com/gofiber/fiber/v2"
)

func SignUpService(c *fiber.Ctx) error {
	data := new(SignUpModel)
	if err := c.BodyParser(data); err != nil {
		return JSONParsingError(c)
	}
	if data.IsValid() != 0 {
		code := data.IsValid()
		switch code {
		case 1:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fiber.Map{
					"ru": "Поле firstName не может быть пустым",
					"en": "Field firstName cannot be empty",
				},
				"code":   400,
				"status": false,
			})
		case 2:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fiber.Map{
					"ru": "Поле LastName не может быть пустым",
					"en": "Field LastName cannot be empty",
				},
				"code":   400,
				"status": false,
			})
		case 3:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fiber.Map{
					"ru": "Поле UserName не может быть пустым",
					"en": "Field UserName cannot be empty",
				},
				"code":   400,
				"status": false,
			})
		case 4:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fiber.Map{
					"ru": "Поле Email не может быть пустым",
					"en": "Field Email cannot be empty",
				},
				"code":   400,
				"status": false,
			})
		case 5:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fiber.Map{
					"ru": "Поле Password не может быть пустым",
					"en": "Field Password cannot be empty",
				},
				"code":   400,
				"status": false,
			})
		}

	}

	if !CheckUniqueUsername(data.UserName) {
		return UniqueUsernameError(c)
	}
	if !CheckEmailUnique(data.Email) {
		return UniqueEmailError(c)
	}

	var profile models.Profile = models.CreateProfile()

	db := global.GetDataBase() // Получение экземпляра базы данных
	transaction := db.Begin()  // Начало транзакции базы данных

	if err := transaction.Save(&profile).Error; err != nil {
		// Сохранение профиля в базе данных. Если происходит ошибка, откатываем транзакцию и возвращаем ошибку
		transaction.Rollback()
		return TransactionError(c)
	} else {
		var rating models.RatingAll = models.RatingAll{Score: 0}
		if err := transaction.Save(&rating).Error; err != nil {
			// Создание и сохранение рейтинга пользователя. Если происходит ошибка, откатываем транзакцию и возвращаем ошибку
			transaction.Rollback()
			return TransactionError(c)
		} else {
			var ratingWeek models.RatingWeek = models.RatingWeek{Score: 0}
			if err := transaction.Save(&ratingWeek).Error; err != nil {
				// Создание и сохранение недельного рейтинга пользователя. Если происходит ошибка, откатываем транзакцию и возвращаем ошибку
				transaction.Rollback()
				return TransactionError(c)
			} else {
				var ratingMonth models.RatingMonth = models.RatingMonth{Score: 0}
				if err := transaction.Save(&ratingMonth).Error; err != nil {
					// Создание и сохранение месячного рейтинга пользователя. Если происходит ошибка, откатываем транзакцию и возвращаем ошибку
					transaction.Rollback()
					return TransactionError(c)
				} else {
					var emailValidation models.EmailValidation = models.GenerateEmailValidation()
					if err := transaction.Save(&emailValidation).Error; err != nil {
						// Генерация и сохранение данных для проверки электронной почты пользователя. Если происходит ошибка, откатываем транзакцию и возвращаем ошибку
						transaction.Rollback()
						return TransactionError(c)
					} else {
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

						if err := transaction.Save(&user).Error; err != nil {
							// Создание и сохранение нового пользователя со всеми связанными данными. Если происходит ошибка, откатываем транзакцию и возвращаем ошибку
							transaction.Rollback()
							return TransactionError(c)
						} else {
							transaction.Commit()
							// Фиксируем транзакцию и возвращаем успешный статус в формате JSON
							return c.Status(200).JSON(fiber.Map{
								"status": true,
							})
						}
					}
				}
			}
		}
	}
}

func SignInService(c *fiber.Ctx) error {
	data := new(SignInModel)
	if err := c.BodyParser(data); err != nil {
		return JSONParsingError(c)
	}
	if data.IsValid() != 0 {
		code := data.IsValid()
		switch code {
		case 1:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fiber.Map{
					"ru": "Поле Email не может быть пустым",
					"en": "Field Email cannot be empty",
				},
				"code":   400,
				"status": false,
			})
		case 2:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fiber.Map{
					"ru": "Поле Password не может быть пустым",
					"en": "Field Password cannot be empty",
				},
				"code":   400,
				"status": false,
			})
		}
	}

	var user models.User

	db := global.GetDataBase() // Получение экземпляра базы данных
	if err := db.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return UserNotFoundByEmailError(c)
	}

	if !user.CheckPassword(data.Password) {
		return IncorrectPasswordError(c)
	}

	token, _ := global.CreateToken(int(user.ID))

	return c.Status(200).JSON(fiber.Map{
		"status": true,
		"token":  token,
	})
}
