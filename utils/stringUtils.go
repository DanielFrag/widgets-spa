package utils

import (
	"math/rand"
	"time"
)

//GenerateRandomAlphaNumericString create a randomic alpha-numeric string
func GenerateRandomAlphaNumericString(length int) string {
	rand.Seed(time.Now().UnixNano())
	numbersAndLetters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if length < 1 {
		return ""
	}
	result := make([]rune, length)
	for i := 0; i < length; i++ {
		result[i] = numbersAndLetters[rand.Intn(len(numbersAndLetters))]
	}
	return string(result)
}
