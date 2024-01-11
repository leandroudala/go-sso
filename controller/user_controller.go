package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"udala/sso/exception"
	"udala/sso/helper"
	"udala/sso/model"
	"udala/sso/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(db *gorm.DB) *UserController {
	service := service.NewUserService(db)

	return &UserController{service: service}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var userForm model.UserFormDTO
	if err := c.ShouldBindJSON(&userForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// creating a new user
	newUser, except := u.service.Create(&userForm)

	if except.HasError() {
		except.Abort(c)
		return
	}

	c.JSON(http.StatusCreated, newUser.ToDTO())
}

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.service.GetAll()

	if err != nil {
		log.Fatal(err)
		except := exception.ApplicationException{
			StatusCode: 400,
			Message:    "Error while retrieving list of users",
		}
		except.Abort(c)
		return
	}

	c.JSON(http.StatusOK, &users)
}

func (u *UserController) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, except := helper.StringToUint64(idParam)

	if except.HasError() {
		except.Abort(c)
		return
	}

	user, except := u.service.GetUserByID(id)

	if except.HasError() {
		except.Abort(c)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, except := helper.StringToUint64(idParam)

	if except.HasError() {
		except.Abort(c)
		return
	}

	except = u.service.DeleteUserById(id)

	if except.HasError() {
		except.Abort(c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Checks if a username or an email is available
func (this *UserController) CheckAvailability(c *gin.Context) {
	email := c.Query("email")
	username := c.Query("username")

	availability, except := this.service.CheckAvailability(email, username)
	if except.HasError() {
		except.Abort(c)
		return
	}

	c.JSON(http.StatusOK, availability)
}
