package generator

import (
	"math/rand"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

type (
	PasswordGenerator interface {
		GeneratePassword() string
		GenerateSeedPhrasePassword() string
	}

	passwordGenerator struct {
		dictionary []string
	}
)

func NewPasswordGenerator() PasswordGenerator {
	return &passwordGenerator{
		dictionary: bip39.GetWordList(),
	}
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

func (pg *passwordGenerator) GenerateSeedPhrasePassword() string {

	length := 12

	words := make([]string, length)

	for i := range words {
		words[i] = pg.dictionary[rand.Intn(len(pg.dictionary))]
	}

	return strings.Join(words, "-")
}
