package app

import (
	"fmt"

	"github.com/Leonid-Yakovlev63/password-generator/internal/generator"
)

func Run(count int, mode string) {

	passwordGenerator := generator.NewPasswordGenerator()

	for i := 0; i < count; i++ {
		switch mode {
		case "random":
			fmt.Println(passwordGenerator.GeneratePassword())
		case "seed":
			fmt.Println(passwordGenerator.GenerateSeedPhrasePassword())
		default:
			fmt.Println(passwordGenerator.GeneratePassword())
		}
	}
}
