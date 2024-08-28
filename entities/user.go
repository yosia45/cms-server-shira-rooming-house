package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;uniqueIndex"`
	FullName string `json:"full_name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;uniqueIndex"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`
}
type RegisterUserRequest struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTUserPayload struct {
	UserID   uint
	Username string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	u.Password = string(hasedPassword)
	return
}
