package controller

import (
	"net/http"
	"udala/sso/model"
	"udala/sso/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{service: *service.NewAuthService(db)}
}

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
