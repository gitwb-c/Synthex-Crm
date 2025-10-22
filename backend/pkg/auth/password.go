package auth

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidatePasswordLenght(password string) (bool, error) {
	passwordMaxLenStr := os.Getenv("EMPLOYEEAUTH_PASSWORD_MAXLEN")
	if len(passwordMaxLenStr) == 0 {
		return false, fmt.Errorf("EMPLOYEEAUTH_PASSWORD_MAXLEN not set")
	}

	passwordMaxLen, err := strconv.Atoi(passwordMaxLenStr)
	if err != nil {
		return false, fmt.Errorf("invalid EMPLOYEEAUTH_PASSWORD_MAXLEN value, must be an integer")
	}

	if len(password) == 0 || len(password) > passwordMaxLen {
		return false, fmt.Errorf("invalid password length")
	}

	return true, nil
}
