package services

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID   string `json:"user_id"`
	UserRole int    `json:"user_role"`
}
