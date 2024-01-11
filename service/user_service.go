package service

import (
	"log"
	"strings"
	"udala/sso/dto"
	"udala/sso/exception"
	"udala/sso/handler"
	"udala/sso/helper"
	"udala/sso/model"
	"udala/sso/repository"

	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	repo := repository.NewUserRepository(db)
	return &UserService{repo: repo}
}

func (service *UserService) Create(dto *model.UserFormDTO) (*model.User, exception.ApplicationException) {
	var user model.User

	hashedPassword, err := helper.HashPassword(dto.Password)
	if err != nil {
		log.Println(err.Error())
		return &user, exception.ApplicationException{
			StatusCode: 400,
			Message:    "Error while encrypting password",
		}
	}

	user = model.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: hashedPassword,
		Name:     dto.Name,
	}

	log.Println("Creating user")
	if err := service.repo.Create(&user); err != nil {
		return &user, handleUserError(err)

	}

	return &user, exception.NilError()
}

func handleUserError(err error) exception.ApplicationException {
	message := err.Error()
	if strings.Contains(message, "Error 1062") {
		return exception.UserDuplicatedException()
	}

	log.Println(err.Error())
	return exception.ApplicationException{
		StatusCode: 400,
		Message:    "Unknown error while creating a new user.",
	}
}

func (service *UserService) GetAll() ([]*model.UserDTO, error) {
	var usersDTO = []*model.UserDTO{}

	users, err := service.repo.GetAllActive()

	if err != nil {
		return usersDTO, err
	}

	for _, user := range users {
		usersDTO = append(usersDTO, user.ToDTO())
	}

	return usersDTO, nil
}

func (service *UserService) GetUserByID(id uint64) (*model.UserDTO, exception.ApplicationException) {
	user, err := service.repo.FindById(id)

	if err != nil {
		log.Fatalln(err)
		return nil, exception.ApplicationException{
			StatusCode: 400,
			Message:    "Error while finding user",
		}
	}

	if user == nil {
		return nil, exception.UserNotFoundException(id)
	}

	return user.ToDTO(), exception.NilError()
}

func (service *UserService) DeleteUserById(id uint64) exception.ApplicationException {
	user, err := service.repo.FindById(id)

	if err != nil {
		log.Fatalln(err.Error())
		return exception.BadRequestException("Error while removing user: could not find user id")
	}

	if user == nil {
		return exception.UserNotFoundException(id)
	}

	err = service.repo.Delete(user)

	if err != nil {
		log.Fatalln(err.Error())
		return exception.BadRequestException("Error while removing user: could not delete user")
	}

	return exception.NilError()
}

func (service *UserService) UserFindByLoginAndPassword(login string, password string) (*model.User, exception.ApplicationException) {

	user, err := service.FindUserByLogin(login)

	if err != nil {
		return user, exception.InvalidLoginDataException()
	}

	validPassword := helper.CheckPassword(password, user.Password)

	if !validPassword {
		return &model.User{}, exception.InvalidLoginDataException()

	}

	return user, exception.NilError()
}

func (service *UserService) FindUserByLogin(login string) (*model.User, error) {
	// check if it is an email
	if strings.Contains(login, "@") {
		return service.FindUserByEmail(login)
	}

	return service.FindUserByUsername(login)
}

func (service *UserService) FindUserByEmail(email string) (*model.User, error) {
	return service.repo.FindByEmail(email)
}

func (service *UserService) FindUserByUsername(username string) (*model.User, error) {
	return service.repo.FindByUsername(username)
}

func (this *UserService) CheckAvailability(email string, username string) (*dto.UserAvailabilityDTO, exception.ApplicationException) {
	email = strings.TrimSpace(email)
	username = strings.TrimSpace(username)

	availability := dto.UserAvailabilityDTO{
		EmailAvailable:    true,
		UsernameAvailable: true,
	}

	// return BadRequest if email and username are not informed
	if email == "" && username == "" {
		return &availability, exception.BadRequestException("Please provide a username or an email to check availability")

	}

	// check if email is duplicated
	if email != "" {
		user, err := this.repo.FindByEmail(email)

		isNotFound := handler.IsNotFoundError(err)
		if err != nil && !isNotFound {
			log.Println(err.Error())
			return &availability, exception.InternalServerError("Could not check if email exists")
		}

		availability.EmailAvailable = user == nil
	}

	// check if username is duplicated
	if username != "" {
		user, err := this.repo.FindByUsername(username)

		isNotFound := handler.IsNotFoundError(err)
		if err != nil && !isNotFound {
			log.Println("aqui ", err.Error())
			return &availability, exception.InternalServerError("Could not check if username exists")
		}

		availability.UsernameAvailable = user == nil
	}

	return &availability, exception.NilError()
}
