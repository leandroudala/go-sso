package model

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type LoginDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type JWTClaims struct {
	UserName string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type JWTToken struct {
	Type  string `json:"type" binding:"required"`
	Token string `json:"token" binding:"required"`
}

type ForgetPasswordForm struct {
	Email string `json:"email" binding:"required"`
}

func (dto *LoginDTO) IsValidData() bool {
	dto.Login = strings.TrimSpace(dto.Login)

	password := strings.TrimSpace(dto.Password)

	return dto.Login != "" && password != ""
}
