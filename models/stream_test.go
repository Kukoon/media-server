package models

import (
	"github.com/google/uuid"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamHasPermission(t *testing.T) {
	assert := assert.New(t)
	db := DatabaseForTesting()

	obj, err := Stream{}.HasPermission(db.DB, uuid.Nil, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Stream{}.HasPermission(db.DB, uuid.Nil, TestStreamID1)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Stream{}.HasPermission(db.DB, TestUserID1, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Stream{}.HasPermission(db.DB, TestUserID1, TestStreamID1)
	assert.NoError(err)
	assert.NotNil(obj)
}
