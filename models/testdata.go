package models

import (
	"time"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	loc = time.FixedZone("UTC+2", +2*60*60)

	testdataChannel1          = uuid.MustParse("df1555f5-7046-4f7a-adcc-195b73949723")
	testdataRecording1        = uuid.MustParse("542685cb-3693-e720-a957-f008f5dae3ee")
	testdataRecording1Format1 = uuid.MustParse("3a4f9157-65bf-4d15-a82b-1cd9295d07e0")
	testdataRecording2        = uuid.MustParse("45da89a7-e5e0-5104-b937-6d4c2d4b6d00")
	testdataRecording2Format1 = uuid.MustParse("09858a43-0532-4ad8-8694-92ed71372ff4")
	testdataRecording3        = uuid.MustParse("edb1cfbb-3476-d639-b3f5-795fabf4ef4d")
	testdataRecording3Format1 = uuid.MustParse("6b1b95f2-d92d-4da7-b56c-1ba86ff22dcd")
)

var testdata = []*gormigrate.Migration{
	{
		ID: "10-data-0010-01-channel",
		Migrate: func(tx *gorm.DB) error {
			return tx.Create(&Channel{
				ID:         testdataChannel1,
				CommonName: "kukoon",
				Title:      "Im Kukoon",
				Logo:       "https://media.kukoon.de/static/css/kukoon/logo.png",
				Secret:     uuid.MustParse("1f349cf3-196d-4e39-9d22-3e35497e990c"),
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Delete(&Recording{
				ID: uuid.MustParse("df1555f5-7046-4f7a-adcc-195b73949723"),
			}).Error
		},
	},
	{
		ID: "10-data-0020-01-recording-1",
		Migrate: func(tx *gorm.DB) error {
			err := tx.Create(&Recording{
				ID:         testdataRecording1,
				ChannelID:  testdataChannel1,
				CommonName: "2020-12-polizeigewalt",
				Poster:     "https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251mp4",
				CreatedAt:  time.Date(2020, 12, 10, 20, 0, 0, 0, loc),
				Duration:   time.Hour,
				Public:     true,
				Listed:     true,
			}).Error
			if err != nil {
				return err
			}
			return tx.Create(&RecordingFormat{
				ID:          testdataRecording1Format1,
				RecordingID: testdataRecording1,
				IsVideo:     true,
				URL:         "https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/542685cb-3693-e720-a957-f008f5dae3ee_20201211_165251.mp4",
				Bytes:       3323919713,
				Resolution:  "1920x1080",
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Delete(&RecordingFormat{
				ID: testdataRecording1Format1,
			}).Error
			if err != nil {
				return err
			}
			return tx.Delete(&Recording{
				ID: testdataRecording1,
			}).Error
		},
	},
	{
		ID: "10-data-0020-01-recording-2",
		Migrate: func(tx *gorm.DB) error {
			err := tx.Create(&Recording{
				ID:         testdataRecording2,
				ChannelID:  testdataChannel1,
				CommonName: "2021-01-faschistische_jahrhundert",
				Poster:     "https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00_20210201_111916mp4",
				CreatedAt:  time.Date(2021, 01, 29, 20, 0, 0, 0, loc),
				Duration:   time.Hour,
				Public:     true,
				Listed:     true,
			}).Error
			if err != nil {
				return err
			}
			return tx.Create(&RecordingFormat{
				ID:          testdataRecording2Format1,
				RecordingID: testdataRecording2,
				IsVideo:     true,
				URL:         "https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/45da89a7-e5e0-5104-b937-6d4c2d4b6d00_20210201_111916.mp4",
				Bytes:       2878429977,
				Resolution:  "1920x1080",
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Delete(&RecordingFormat{
				ID: testdataRecording2Format1,
			}).Error
			if err != nil {
				return err
			}
			return tx.Delete(&Recording{
				ID: testdataRecording2,
			}).Error
		},
	},
	{
		ID: "10-data-0020-01-recording-3",
		Migrate: func(tx *gorm.DB) error {
			err := tx.Create(&Recording{
				ID:         testdataRecording3,
				ChannelID:  testdataChannel1,
				CommonName: "2021-02-pushbacks",
				Poster:     "https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d_20210207_111945mp4",
				CreatedAt:  time.Date(2021, 2, 5, 20, 0, 0, 0, loc),
				Duration:   time.Hour,
				Public:     true,
				Listed:     true,
			}).Error
			if err != nil {
				return err
			}
			return tx.Create(&RecordingFormat{
				ID:          testdataRecording3Format1,
				RecordingID: testdataRecording3,
				IsVideo:     true,
				URL:         "https://media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/edb1cfbb-3476-d639-b3f5-795fabf4ef4d_20210207_111945.mp4",
				Bytes:       1092701356,
				Resolution:  "1920x1080",
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Delete(&RecordingFormat{
				ID: testdataRecording3Format1,
			}).Error
			if err != nil {
				return err
			}
			return tx.Delete(&Recording{
				ID: testdataRecording3,
			}).Error
		},
	},
}
