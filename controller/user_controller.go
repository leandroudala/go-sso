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

// @Summary Create user
// @Schemes
// @Description Create a new user
// @ID user-create
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.UserFormDTO true "user info"
// @Success 200 {object} model.UserDTO
// @Router /users [POST]
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

// @Summary Get all users
// @Schemes
// @Description Get all users in the system
// @Tags User
// @ID user-all
// @Produce json
// @Success 200 {array} model.UserDTO
// @Router /users [get]
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

// @Summary Get user by ID
// @Schemes
// @Description Get user by ID
// @Tags User
// @ID user-get
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.UserDTO
// @Failure 400 {object} exception.ApplicationException
// @Router /users/{id} [get]
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

// @Summary Delete user
// @Schemes
// @Tags User
// @Description Delete user
// @ID user-delete
// @Param id path int true "User ID"
// @Success 204 {object} exception.ApplicationException
// @Failure 404 {object} exception.ApplicationException
// @Router /users/{id} [DELETE]
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

// @Summary Check User Availability
// @Schemes
// @Tags User
// @Description Check if email and/or username is in use
// @ID user-availability
// @Param email query string false "Email"
// @Param username query string false "Username"
// @Success 204 {object} exception.ApplicationException
// @Failure 404 {object} exception.ApplicationException
// @Router /users/check-availability [GET]
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
