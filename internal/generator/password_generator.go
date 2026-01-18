package generator

import (
	"math/rand"
)

type (
	PasswordGenerator interface {
		GeneratePassword() string
	}

	passwordGenerator struct{}
)

func NewPasswordGenerator() PasswordGenerator {
	return &passwordGenerator{}
}

func (pg *passwordGenerator) GeneratePassword() string {

	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;':\",./<>?")

	length := rand.Intn(15) + 12

	password := make([]rune, length)

	for i := range password {
		password[i] = runes[rand.Intn(len(runes))]
	}

	return string(password)

}
