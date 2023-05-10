package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateString(n int) string {
	rand.Seed(time.Now().UnixNano())
	m := len(alphabet)

	res := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		pick := rand.Intn(m)
		res = append(res, alphabet[pick])
	}

	return string(res)
}

func GenerateInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
