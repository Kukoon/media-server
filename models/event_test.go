package models

import (
	"github.com/google/uuid"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventHasPermission(t *testing.T) {
	assert := assert.New(t)
	db := DatabaseForTesting()

	obj, err := Event{}.HasPermission(db.DB, uuid.Nil, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Event{}.HasPermission(db.DB, uuid.Nil, TestEventID1)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Event{}.HasPermission(db.DB, TestUserID1, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Event{}.HasPermission(db.DB, TestUserID1, TestEventID1)
	assert.NoError(err)
	assert.NotNil(obj)
}
