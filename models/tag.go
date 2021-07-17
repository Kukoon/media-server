package models

import (
	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Tag struct
type Tag struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"df1555f5-7046-4f7a-adcc-195b73949723"`
	Lang *TagLang  `json:"lang"`
}

// TagLang struct
type TagLang struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"df1555f5-7046-4f7a-adcc-195b73949723"`
	Tag   Tag       `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TagID uuid.UUID `json:"-" gorm:"type:uuid;unique_index:idx_tag_lang"`
	// attributes
	Lang string `json:"lang" gorm:"unique_index:idx_tag_lang" example:"de"`
	Name string `json:"name" example:"Buchvorstellung"`
}

var (
	// TestTagBuchvorstellungID - uuid for tag buchvorstellung
	TestTagBuchvorstellungID = uuid.MustParse("0bca0cf4-a9b9-46d7-821f-18c59c08fc1d")
	// TestTagBuchvorstellungLangID - german translation
	TestTagBuchvorstellungLangID = uuid.MustParse("35822fe2-1910-48e7-904f-15c9e6f7ea34")

	// TestTagDiskussionID - uuid for tag diskussion
	TestTagDiskussionID = uuid.MustParse("277026b0-b9d6-48d6-bfa1-96dcc7eb3451")
	// TestTagDiskussionLangID - german translation
	TestTagDiskussionLangID = uuid.MustParse("38722845-beba-4e3d-ad3f-694c029d751f")

	// TestTagVortragID - uuid for tag vortrag
	TestTagVortragID = uuid.MustParse("7297a654-71f9-43be-8120-69b8152f01fc")
	// TestTagVortragLangID - german translation
	TestTagVortragLangID = uuid.MustParse("ec784c8e-2673-4870-b219-eb636e4765c8")

	// TestTagKonzertID - uuid for tag konzert
	TestTagKonzertID = uuid.MustParse("71082693-c58d-43b7-86ff-6b240c643a83")
	// TestTagKonzertLangID - german translation
	TestTagKonzertLangID = uuid.MustParse("57160845-0982-473c-820b-c0b9a132c282")
)

func init() {
	migrations = append(migrations, []*gormigrate.Migration{
		{
			ID: "01-schema-0018-01-tag",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Tag{},
					&TagLang{})
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("tag_langs"); err != nil {
					return err
				}
				return nil
			},
		},
	}...)

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0018-01-tag-eg",
			Migrate: func(tx *gorm.DB) error {
				// -
				if err := tx.Create(&Tag{
					ID: TestTagBuchvorstellungID,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&TagLang{
					ID:    TestTagBuchvorstellungLangID,
					TagID: TestTagBuchvorstellungID,
					Lang:  "de",
					Name:  "Buchvorstellung",
				}).Error; err != nil {
					return err
				}
				// -
				if err := tx.Create(&Tag{
					ID: TestTagDiskussionID,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&TagLang{
					ID:    TestTagDiskussionLangID,
					TagID: TestTagDiskussionID,
					Lang:  "de",
					Name:  "Diskussion",
				}).Error; err != nil {
					return err
				}
				// -
				if err := tx.Create(&Tag{
					ID: TestTagVortragID,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&TagLang{
					ID:    TestTagVortragLangID,
					TagID: TestTagVortragID,
					Lang:  "de",
					Name:  "Vortrag",
				}).Error; err != nil {
					return err
				}
				// -
				if err := tx.Create(&Tag{
					ID: TestTagKonzertID,
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&TagLang{
					ID:    TestTagKonzertLangID,
					TagID: TestTagKonzertID,
					Lang:  "de",
					Name:  "Konzert",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				// -
				if err := tx.Delete(&TagLang{ID: TestTagKonzertLangID}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Tag{ID: TestTagKonzertID}).Error; err != nil {
					return err
				}
				// -
				if err := tx.Delete(&TagLang{ID: TestTagVortragLangID}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Tag{ID: TestTagVortragID}).Error; err != nil {
					return err
				}
				// -
				if err := tx.Delete(&TagLang{ID: TestTagDiskussionLangID}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Tag{ID: TestTagDiskussionID}).Error; err != nil {
					return err
				}
				// -
				if err := tx.Delete(&TagLang{ID: TestTagBuchvorstellungLangID}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Tag{ID: TestTagBuchvorstellungID}).Error; err != nil {
					return err
				}
				return nil
			},
		},
	}...)
}
