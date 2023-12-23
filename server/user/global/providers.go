package global

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// Функция для создания JWT токена
func CreateToken(user_id int) (string, error) {
	// Создание нового токена
	token := jwt.New(jwt.SigningMethodHS256)

	// Установка данных в токен
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id

	// Генерация подписи токена
	secret := []byte(JWT_SECRET_KEY)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Функция для дешифровки JWT токена
func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	// Парсинг токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	// Проверка валидности токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
