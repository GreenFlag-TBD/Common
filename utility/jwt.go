package utility

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	key_exp = "exp"
	key_id  = "id"
)

func GenerateToken(id string, seed string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		key_exp: time.Now().Add(duration).Unix(),
		key_id:  id,
	})
	tokenString, err := token.SignedString([]byte(seed))
	if err != nil {
		fmt.Printf("\nError while generating token: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string, seed string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(seed), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string), nil
	}

	return "", nil
}

func GenerateNONTTLToken(id string, seed string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		key_id: id,
	})
	tokenString, err := token.SignedString([]byte(seed))
	if err != nil {
		fmt.Printf("\nError while generating token: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
