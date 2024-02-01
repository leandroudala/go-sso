package service

import (
	"os"
	"time"
	"udala/sso/exception"
	"udala/sso/model"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var secretKey = []byte(os.Getenv("SECRET_KEY_JWT"))

type AuthService struct {
	userService UserService
	smtpService SmtpService
}

func NewAuthService(db *gorm.DB) *AuthService {
	userService := NewUserService(db)
	smtpService := NewSmtpService()

	return &AuthService{userService: *userService, smtpService: *smtpService}
}

func (service *AuthService) AuthLogin(dto model.LoginDTO) (string, exception.ApplicationException) {

	var claim model.JWTClaims
	var signedToken string

	// check if data is valid
	if !dto.IsValidData() {
		return signedToken, exception.InvalidLoginDataException()
	}

	// check if user and password match
	user, except := service.userService.UserFindByLoginAndPassword(dto.Login, dto.Password)
	if except.HasError() {
		return signedToken, except
	}

	claim = model.JWTClaims{
		UserName: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "udala.app",
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return signedToken, exception.InvalidLoginDataException()
	}

	return signedToken, exception.NilError()
}
