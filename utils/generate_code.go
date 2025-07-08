package utils

import (
	"math/rand"
	"time"
)

func GenerateCode(length int) string {
	const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	build := make([]byte, length)
	for i := range build{
		build[i] = char[random.Intn(len(char))]
	}
	return  string(build)
}