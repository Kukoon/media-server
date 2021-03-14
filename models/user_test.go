package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserPassword(t *testing.T) {
	assert := assert.New(t)
	password := "password"
	user, err := NewUser("admin", password)

	assert.Nil(err)
	assert.NotNil(user)

	assert.False(user.ValidatePassword("12346"))
	assert.True(user.ValidatePassword(password))
	assert.NotEqual(password, user.Password, "password should be hashed")
}
