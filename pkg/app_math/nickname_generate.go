package app_math

import (
	"math/rand"
	"time"
)

// 生成随机用户名
func GenerateNickname(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	username := make([]rune, length)
	for i := range username {
		username[i] = letters[rand.Intn(len(letters))]
	}
	return string(username)
}
