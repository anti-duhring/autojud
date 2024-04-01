package jwt

import (
	"time"

	"github.com/anti-duhring/goncurrency/pkg/logger"
	"github.com/dgrijalva/jwt-go"
)

var (
	SecretKey = []byte("secret")
)

type JwtClaims struct {
	UserID string  `json:"userId"`
	Exp    float64 `json:"exp"`
}

func GenerateToken(userID string) (string, int64, error) {

	exp := time.Now().Add(time.Hour * 24 * 7).Unix()

	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		logger.Error("Error generating token", err)
		return "", 0, err
	}

	return tokenString, exp, nil
}

func TokenExpired(exp int64) bool {
	return time.Now().Unix() > exp
}

func ParseToken(tokenStr string) (*JwtClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &JwtClaims{
			UserID: claims["userID"].(string),
			Exp:    claims["exp"].(float64),
		}, nil
	}

	return nil, err
}
