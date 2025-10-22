package auth

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

func isValidEmailFormat(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func validateDomain(email string) (bool, error) {
	if !isValidEmailFormat(email) {
		return false, fmt.Errorf("invalid email format")
	}

	domain := strings.Split(email, "@")[1]
	if len(domain) == 0 {
		return false, fmt.Errorf("invalid email: missing domain")
	}

	_, err := net.LookupMX(domain)
	if err != nil {
		return false, fmt.Errorf("domain does not have an MX record: %v", err)
	}

	return true, nil
}

func ValidateEmail(email string) (bool, error) {
	valid, err := validateDomain(email)
	if !valid {
		return false, err
	}

	return true, nil
}
