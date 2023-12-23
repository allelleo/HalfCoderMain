package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateCode(length int) string {
	// Define the characters that can be used in the password
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789----------"

	// Create a buffer to store the generated password
	password := make([]byte, length)

	// Generate random indexes to select characters from the charset
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		password[i] = charset[index.Int64()]
	}

	return string(password)
}
