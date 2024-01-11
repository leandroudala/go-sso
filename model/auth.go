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
	UserID uint64 `json:"userId"`
	jwt.StandardClaims
}

type JWTToken struct {
	Type  string `json:"type" binding:"required"`
	Token string `json:"token" binding:"required"`
}

func (dto *LoginDTO) IsValidData() bool {
	dto.Login = strings.TrimSpace(dto.Login)

	password := strings.TrimSpace(dto.Password)

	return dto.Login != "" && password != ""
}
