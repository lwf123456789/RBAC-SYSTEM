package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
