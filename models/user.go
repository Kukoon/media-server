package models

import (
	"dev.sum7.eu/genofire/golang-lib/web/auth"
	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"88078ec0-2135-445f-bf05-632701c77695"`
	Username   string     `json:"username" gorm:"unique" example:"kukoon"`
	Password   string     `json:"-" example:"super secret password"`
	ForgetCode *uuid.UUID `json:"-" gorm:"forget_code;type:uuid"`
	Channels   []Channel  `json:"-" gorm:"many2many:user_channels;"`
}

var (
	// TestUserID1 - uuid of initial user
	TestUserID1 = uuid.MustParse("88078ec0-2135-445f-bf05-632701c77695")
	TestUserID2 = uuid.MustParse("aca3f761-042b-4c7f-a8d1-82e972683adf")
)

func init() {
	migrations = append(migrations, []*gormigrate.Migration{
		{
			ID: "01-schema-0008-01-user",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		// there should be an init user
		{
			ID: "10-data-0008-01-user",
			Migrate: func(tx *gorm.DB) error {
				user, err := auth.NewUser("kukoon", "CHANGEME")
				if err != nil {
					return err
				}
				user.ID = TestUserID1
				return tx.Create(user).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Delete(&User{
					ID: TestUserID1,
				}).Error
			},
		},
	}...)

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0008-02-user",
			Migrate: func(tx *gorm.DB) error {
				user, err := auth.NewUser("c3woc", "CHANGEME")
				if err != nil {
					return err
				}
				user.ID = TestUserID2
				return tx.Create(user).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Delete(&User{
					ID: TestUserID2,
				}).Error
			},
		},
	}...)
}
