package jwtpkg

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gitwb-c/crm.saas/backend/internal/db/cache"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type TokenClaims struct {
	SessionId string
}

func GenerateToken() (string, string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	sessionId, err := cache.GenerateSessionID()
	if err != nil {
		return "", "", err
	}
	claims["sessionId"] = sessionId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenStr, err := token.SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}
	return tokenStr, sessionId, nil
}

func ParseToken(tokenStr string) (TokenClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return TokenClaims{}, fmt.Errorf("invalid token or missing claims")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sessionId := claims["sessionId"].(string)
		return TokenClaims{
			SessionId: sessionId,
		}, nil
	}
	return TokenClaims{}, fmt.Errorf("invalid token or missing claims")
}
