package generator_test

import (
	"fmt"
	"testing"

	"github.com/Leonid-Yakovlev63/password-generator/internal/generator"
)

func setupPasswordGenerator(t *testing.T) generator.PasswordGenerator {
	return generator.NewPasswordGenerator()
}

func TestPasswordGenerator(t *testing.T) {
	testGenerator := setupPasswordGenerator(t)

	for range 100 {
		password := testGenerator.GeneratePassword()
		fmt.Println(password)
		if len(password) < 12 {
			t.Errorf("Длина пароля меньше минимальной: %d символов (ожидается >= 12)", len(password))
		}
	}
}
