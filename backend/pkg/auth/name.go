package auth

import (
	"fmt"
	"os"
	"strconv"
)

func ValidateNameLenght(name string) (bool, error) {
	nameMaxLenStr := os.Getenv("EMPLOYEEAUTH_PASSWORD_MAXLEN")
	nameMinLenStr := os.Getenv("EMPLOYEEAUTH_PASSWORD_MINLEN")

	if len(nameMinLenStr) == 0 || len(nameMaxLenStr) == 0 {
		return false, fmt.Errorf("EMPLOYEEAUTH_PASSWORD_MAXLEN or EMPLOYEEAUTH_PASSWORD_MINLEN not set")
	}

	nameMaxLen, errMax := strconv.Atoi(nameMaxLenStr)
	nameMinLen, errMin := strconv.Atoi(nameMinLenStr)

	if errMax != nil || errMin != nil {
		return false, fmt.Errorf("invalid name length values, must be integers")
	}

	if len(name) < nameMinLen || len(name) > nameMaxLen {
		return false, fmt.Errorf("name length must be between %d and %d characters", nameMinLen, nameMaxLen)
	}

	return true, nil

}
