package service

import (
	"time"
	"udala/sso/exception"
	"udala/sso/model"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var secretKey = []byte("SECRET_KEY_HERE_OKAY?")

type AuthService struct {
	userService UserService
}

func NewAuthService(db *gorm.DB) *AuthService {
	userService := NewUserService(db)

	return &AuthService{userService: *userService}
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
		UserID: user.ID,
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
