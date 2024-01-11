package model

import "time"

// User model
type User struct {
	ID        uint64    `json:"ID" gorm:"primaryKey"`
	Username  string    `json:"Username" gorm:"not null; unique"`
	Name      string    `json:"Name" gorm:"not null"`
	Email     string    `json:"Email" gorm:"not null; unique"`
	Password  string    `json:"Password" gorm:"not null"`
	CreatedAt time.Time `json:"CreatedAt" gorm:"autoCreateTime; not null"`
	Disabled  bool      `json:"Disabled" gorm:"default:false; not null"`
}

func (user *User) ToDTO() *UserDTO {
	return &UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}
}

// DTO to show user
type UserDTO struct {
	ID       uint64 `json:"ID" binding:"required"`
	Username string `json:"Username" binding:"required"`
	Name     string `json:"Name" binding:"required"`
	Email    string `json:"Email" binding:"required,email"`
}

// DTO for new users form
type UserFormDTO struct {
	Username string `json:"Username" binding:"required"`
	Name     string `json:"Name" binding:"required"`
	Email    string `json:"Email" binding:"required,email"`
	Password string `json:"Password" binding:"required"`
}
