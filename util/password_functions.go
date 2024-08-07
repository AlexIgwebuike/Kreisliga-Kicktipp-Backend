package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	cost := 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	hashedPassword := string(bytes)

	return hashedPassword, err
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
