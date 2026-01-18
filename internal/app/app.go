package app

import (
	"fmt"

	"github.com/Leonid-Yakovlev63/password-generator/internal/generator"
)

func Run() {
	passwordGenerator := generator.NewPasswordGenerator()

	fmt.Println(passwordGenerator.GeneratePassword())
}
