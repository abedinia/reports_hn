package tests

import (
	"github.com/dgrijalva/jwt-go"
	jwt_token "report_hn/internal/token"
	"testing"
	"time"
)

func TestMakeToken(t *testing.T) {
	id := uint(12345)

	tokenString, err := jwt_token.MakeToken(id)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return jwt_token.SecretKey, nil
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(float64); !ok || uint(userID) != id {
			t.Errorf("expected user_id %v, got %v", id, userID)
		}
	} else {
		t.Errorf("expected valid token, got %v", err)
	}
}

func TestDecodeToken_InvalidToken(t *testing.T) {
	invalidTokenString := "invalid.token.string"

	_, err := jwt_token.DecodeToken(invalidTokenString)
	if err == nil {
		t.Errorf("expected error, got none")
	}
}

func TestDecodeToken_ExpiredToken(t *testing.T) {
	id := uint(12345)
	expiredTime := time.Now().Add(-time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     expiredTime,
	})
	tokenString, err := token.SignedString(jwt_token.SecretKey)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = jwt_token.DecodeToken(tokenString)
	if err == nil {
		t.Errorf("expected error, got none")
	}
}
