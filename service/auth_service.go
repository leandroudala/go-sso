package service

import (
	"bytes"
	"html/template"
	"log"
	"math/rand"
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
	smtp        SmtpService
}

func NewAuthService(db *gorm.DB) *AuthService {
	userService := NewUserService(db)
	smtpService := NewSmtpService()

	return &AuthService{userService: *userService, smtp: *smtpService}
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

func (service *AuthService) ForgetPassword(form model.SendEmailForm) {

	err := service.smtp.sendEmail(
		form.Email,
		"Redefinição de senha - Udala.app",
		"Acesse este link para recuperar a senha.",
	)

	if err.HasError() {
		log.Println("ForgetPassword SMTP error: " + err.Message)
	}

}

var validCodes = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func (service *AuthService) generateVerificationCode(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, length)

	for i := 0; i < length; i++ {
		index := r.Intn(len(validCodes))
		code[i] = validCodes[index]
	}

	return string(code)
}

var validateEmailTemplate = (`
	Olá!

	Segue o código para confirmar o e-mail: {{.Code}}
`)

func (service *AuthService) ValidateEmail(email string) exception.ApplicationException {
	code := service.generateVerificationCode(6)

	bodyTemplate, err := template.New("validateEmailTemplate").Parse(validateEmailTemplate)

	if err != nil {
		log.Println("Error while converting template", err.Error())
		return exception.InternalServerError("Could not send email to validate email.")
	}

	data := struct {
		Code string
	}{
		Code: code,
	}

	var buf bytes.Buffer
	err = bodyTemplate.Execute(&buf, data)

	if err != nil {
		log.Println("Error while executing template:", err)
		return exception.InternalServerError("Error while executing template")
	}

	bodyHtml := buf.String()

	service.smtp.sendEmail(
		email,
		"Confirmar Email",
		bodyHtml,
	)

	return exception.NilError()
}
