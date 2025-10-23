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
	passwordMinLenStr := os.Getenv("EMPLOYEEAUTH_PASSWORD_MINLEN")

	if len(passwordMinLenStr) == 0 || len(passwordMaxLenStr) == 0 {
		return false, fmt.Errorf("EMPLOYEEAUTH_PASSWORD_MAXLEN or EMPLOYEEAUTH_PASSWORD_MINLEN not set")
	}

	passwordMaxLen, errMax := strconv.Atoi(passwordMaxLenStr)
	passwordMinLen, errMin := strconv.Atoi(passwordMinLenStr)

	if errMax != nil || errMin != nil {
		return false, fmt.Errorf("invalid password length values, must be integers")
	}

	if len(password) < passwordMinLen || len(password) > passwordMaxLen {
		return false, fmt.Errorf("password length must be between %d and %d characters", passwordMinLen, passwordMaxLen)
	}

	return true, nil
}
