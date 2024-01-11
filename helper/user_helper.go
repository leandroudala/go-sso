package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Use GenerateFromPassword with a suitable cost (e.g., 12)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
