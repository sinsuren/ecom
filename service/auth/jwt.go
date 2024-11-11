package auth

import (
	"ecom/config"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
