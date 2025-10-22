package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type TokenClaims struct {
	EmployeeId   string
	Company      string
	DepartmentID string
}

func GenerateToken(employeeId string, company string, departmentId string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["employeeId"] = employeeId
	claims["company"] = company
	claims["departmentId"] = departmentId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenStr, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ParseToken(tokenStr string) (TokenClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return TokenClaims{}, fmt.Errorf("invalid token or missing claims")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		employeeId := claims["employeeId"].(string)
		company := claims["company"].(string)
		departmentId := claims["departmentId"].(string)
		return TokenClaims{EmployeeId: employeeId,
			Company:      company,
			DepartmentID: departmentId}, nil
	}
	return TokenClaims{}, fmt.Errorf("invalid token or missing claims")
}
