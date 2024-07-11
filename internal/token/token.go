package token

import (
	"github.com/dgrijalva/jwt-go"
	"report_hn/internal/config"
	"time"
)

var SecretKey = []byte(config.AppConfig.App.Token)

func MakeToken(id uint) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(SecretKey)

	return tokenString, err
}

func DecodeToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return SecretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return 0, jwt.NewValidationError("user_id not found in token claims", jwt.ValidationErrorClaimsInvalid)
		}
		return uint(userID), nil
	}

	return 0, jwt.NewValidationError("invalid token", jwt.ValidationErrorSignatureInvalid)
}
