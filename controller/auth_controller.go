package controller

import (
	"net/http"
	"udala/sso/model"
	"udala/sso/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	service     service.AuthService
	smtpService service.SmtpService
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		service:     *service.NewAuthService(db),
		smtpService: *service.NewSmtpService(),
	}
}

// @Summary User Log-on
// @Schemes
// @Tags Logon
// @Description Generates JWT Token for user
// @ID user-logon
// @Param login body model.LoginDTO true "User Login information"
// @Success 200 {object} model.JWTToken
// @Failure 400 {object} exception.ApplicationException
// @Router /auth [POST]
func (con *AuthController) AuthLogin(c *gin.Context) {
	var loginForm model.LoginDTO
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	signedToken, except := con.service.AuthLogin(loginForm)

	if except.HasError() {
		except.Abort(c)
		return
	}

	c.JSON(http.StatusOK, model.JWTToken{
		Type:  "Bearer",
		Token: signedToken,
	})
}

// @Summary User Forget Password
// @Schemes
// @Tags ForgetPassword
// @Description Sends email redefinition when user forgets password
// @ID forget-password
// @Param passwordForm body model.ForgetPasswordForm true "User Email information"
// @Success 204 {object} nil
// @Failure 400 {object} exception.ApplicationException
// @Router /auth/forget-password [POST]
func (con *AuthController) ForgetPassword(c *gin.Context) {
	var forgetForm model.ForgetPasswordForm
	if err := c.ShouldBindJSON(&forgetForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go con.service.ForgetPassword(forgetForm)

	c.Status(http.StatusNoContent)
}
