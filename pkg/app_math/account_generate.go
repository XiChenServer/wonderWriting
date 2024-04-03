package app_math

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "0123456789"
	var result string
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		result += string(charset[randomIndex])
	}
	return result
}
