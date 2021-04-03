package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"88078ec0-2135-445f-bf05-632701c77695"`
	Username string    `json:"username" gorm:"unique" example:"kukoon"`
	Password string    `json:"-" example:"super secret password"`
	Channels []Channel `json:"-" gorm:"many2many:user_channels;"`
}

func NewUser(username, password string) (*User, error) {
	user := &User{
		Username: username,
	}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}
	return user, nil
}

func (this *User) SetPassword(password string) error {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	this.Password = string(p)
	return nil
}

func (this *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(password))
	return err == nil
}
