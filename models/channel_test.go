package models

import (
	"github.com/google/uuid"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelHasPermission(t *testing.T) {
	assert := assert.New(t)
	db := DatabaseForTesting()

	obj, err := Channel{}.HasPermission(db.DB, uuid.Nil, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Channel{}.HasPermission(db.DB, uuid.Nil, TestChannelID1)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Channel{}.HasPermission(db.DB, TestUserID1, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Channel{}.HasPermission(db.DB, TestUserID1, TestChannelID1)
	assert.NoError(err)
	assert.NotNil(obj)
}
