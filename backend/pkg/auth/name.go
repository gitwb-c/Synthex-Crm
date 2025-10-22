package auth

import (
	"fmt"
	"os"
	"strconv"
)

func ValidateNameLenght(name string) (bool, error) {
	nameMaxLenStr := os.Getenv("EMPLOYEEAUTH_PASSWORD_MAXLEN")
	if len(nameMaxLenStr) == 0 {
		return false, fmt.Errorf("EMPLOYEEAUTH_PASSWORD_MAXLEN not set")
	}

	nameMaxLen, err := strconv.Atoi(nameMaxLenStr)
	if err != nil {
		return false, fmt.Errorf("invalid EMPLOYEEAUTH_PASSWORD_MAXLEN value, must be an integer")
	}

	if len(name) == 0 || len(name) > nameMaxLen {
		return false, fmt.Errorf("invalid name length")
	}

	return true, nil
}
