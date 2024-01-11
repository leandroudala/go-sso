package repository

import (
	"udala/sso/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(user *model.User) error {
	return repo.db.Create(&user).Error
}

func (repo *UserRepository) GetAllActive() ([]model.User, error) {
	var users []model.User
	err := repo.db.Where("disabled = ?", false).Find(&users).Error

	return users, err
}

func (repo *UserRepository) FindById(id uint64) (*model.User, error) {
	var user model.User
	err := repo.db.Where("id = ? and disabled = ?", id, false).Find(&user).Error
	if user.ID == 0 {
		return nil, err
	}

	return &user, err
}

func (repo *UserRepository) Delete(user *model.User) error {
	return repo.db.Model(&user).Update("disabled", true).Error
}

func (repo *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := repo.db.Model(&model.User{}).First(&user, "email = ?", email).Error

	if user.ID == 0 {
		return nil, err
	}

	return &user, err
}

func (repo *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := repo.db.Model(&model.User{}).First(&user, "username = ?", username).Error

	if user.ID == 0 {
		return nil, err
	}

	return &user, err
}
