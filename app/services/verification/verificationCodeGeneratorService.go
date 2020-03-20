package verification

import (
	"math/rand"
)

// GenerateVerificationCode is to generate verification code for user
func GenerateVerificationCode(digits int) string{
	var letter = []rune("0123456789")
	b := make([]rune, digits)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}