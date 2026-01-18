package app

import (
	"fmt"

	"github.com/Leonid-Yakovlev63/password-generator/internal/generator"
)

func Run(count int) {

	passwordGenerator := generator.NewPasswordGenerator()

	for i := 0; i < count; i++ {
		fmt.Println(passwordGenerator.GeneratePassword())
	}
}
