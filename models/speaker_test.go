package models

import (
	"github.com/google/uuid"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpeakerHasPermission(t *testing.T) {
	assert := assert.New(t)
	db := DatabaseForTesting()

	obj, err := Speaker{}.HasPermission(db.DB, uuid.Nil, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Speaker{}.HasPermission(db.DB, uuid.Nil, TestStream1IDSpeaker1)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Speaker{}.HasPermission(db.DB, TestUserID1, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Speaker{}.HasPermission(db.DB, TestUserID1, TestStream1IDSpeaker1)
	assert.NoError(err)
	assert.NotNil(obj)
}
