package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

func CheckPassword(hashPass, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))

	if err != nil {
		return false
	}

	return true
}
