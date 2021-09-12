package models

import (
	"github.com/google/uuid"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordingHasPermission(t *testing.T) {
	assert := assert.New(t)
	db := DatabaseForTesting()

	obj, err := Recording{}.HasPermission(db.DB, uuid.Nil, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Recording{}.HasPermission(db.DB, uuid.Nil, TestRecording1ID)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Recording{}.HasPermission(db.DB, TestUserID1, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = Recording{}.HasPermission(db.DB, TestUserID1, TestRecording1ID)
	assert.NoError(err)
	assert.NotNil(obj)
}

func TestRecordingFormatHasPermission(t *testing.T) {
	assert := assert.New(t)
	db := DatabaseForTesting()

	obj, err := RecordingFormat{}.HasPermission(db.DB, uuid.Nil, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = RecordingFormat{}.HasPermission(db.DB, uuid.Nil, TestRecording1IDFormat1)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = RecordingFormat{}.HasPermission(db.DB, TestUserID1, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = RecordingFormat{}.HasPermission(db.DB, TestUserID1, TestRecording1IDFormat1)
	assert.NoError(err)
	assert.NotNil(obj)
}

func TestRecordingLangHasPermission(t *testing.T) {
	assert := assert.New(t)
	db := DatabaseForTesting()

	obj, err := RecordingLang{}.HasPermission(db.DB, uuid.Nil, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = RecordingLang{}.HasPermission(db.DB, uuid.Nil, TestRecording1IDLang1)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = RecordingLang{}.HasPermission(db.DB, TestUserID1, uuid.Nil)
	assert.NoError(err)
	assert.Nil(obj)

	obj, err = RecordingLang{}.HasPermission(db.DB, TestUserID1, TestRecording1IDLang1)
	assert.NoError(err)
	assert.NotNil(obj)
}
