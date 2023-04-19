package helper

import "math/rand"

// define the given charset, char only
var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func CreateRandCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
