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

	// TestTagDiskussionID - uuid for tag diskussion
	TestTagDiskussionID = uuid.MustParse("277026b0-b9d6-48d6-bfa1-96dcc7eb3451")

	// TestTagVortragID - uuid for tag vortrag
	TestTagVortragID = uuid.MustParse("7297a654-71f9-43be-8120-69b8152f01fc")

	// TestTagKonzertID - uuid for tag konzert
	TestTagKonzertID = uuid.MustParse("71082693-c58d-43b7-86ff-6b240c643a83")

	// TestTagAusstellungID - uuid for tag galerie
	TestTagAusstellungID = uuid.MustParse("87fb7f4b-1e5c-4c1b-8020-2860987da6bc")

	// TestTagInterviewID - uuid for tag interview
	TestTagInterviewID = uuid.MustParse("1aa3f441-4461-42da-a858-63abf5ee254c")
)

func init() {
	testTagBuchvorstellungLangDEID := uuid.MustParse("35822fe2-1910-48e7-904f-15c9e6f7ea34")
	testTagDiskussionLangDEID := uuid.MustParse("38722845-beba-4e3d-ad3f-694c029d751f")
	testTagVortragLangDEID := uuid.MustParse("ec784c8e-2673-4870-b219-eb636e4765c8")
	testTagKonzertLangDEID := uuid.MustParse("57160845-0982-473c-820b-c0b9a132c282")
	testTagAusstellungLangDEID := uuid.MustParse("4dc67d8a-be5c-403c-904c-106a5dd83627")
	testTagInterviewLangDEID := uuid.MustParse("366de5e6-fc00-415c-9ef3-bfc8990841ad")

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
				if err := tx.Migrator().DropTable("tags"); err != nil {
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
				if err := tx.Create(&[]*Tag{
					{ID: TestTagBuchvorstellungID},
					{ID: TestTagDiskussionID},
					{ID: TestTagVortragID},
					{ID: TestTagKonzertID},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&[]*TagLang{
					{
						ID:    testTagBuchvorstellungLangDEID,
						TagID: TestTagBuchvorstellungID,
						Lang:  "de",
						Name:  "Buchvorstellung",
					},
					{
						ID:    testTagDiskussionLangDEID,
						TagID: TestTagDiskussionID,
						Lang:  "de",
						Name:  "Diskussion",
					},
					{
						ID:    testTagVortragLangDEID,
						TagID: TestTagVortragID,
						Lang:  "de",
						Name:  "Vortrag",
					},
					{
						ID:    testTagKonzertLangDEID,
						TagID: TestTagKonzertID,
						Lang:  "de",
						Name:  "Konzert",
					},
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&[]*TagLang{
					{ID: testTagKonzertLangDEID},
					{ID: testTagVortragLangDEID},
					{ID: testTagDiskussionLangDEID},
					{ID: testTagBuchvorstellungLangDEID},
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&[]*Tag{
					{ID: TestTagKonzertID},
					{ID: TestTagVortragID},
					{ID: TestTagDiskussionID},
					{ID: TestTagBuchvorstellungID},
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0018-01-tag-eg2",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&[]*Tag{
					{ID: TestTagAusstellungID},
					{ID: TestTagInterviewID},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&[]*TagLang{
					{
						ID:    testTagAusstellungLangDEID,
						TagID: TestTagAusstellungID,
						Lang:  "de",
						Name:  "Ausstellung",
					},
					{
						ID:    testTagInterviewLangDEID,
						TagID: TestTagInterviewID,
						Lang:  "de",
						Name:  "Interview",
					},
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&[]*TagLang{
					{ID: testTagInterviewLangDEID},
					{ID: testTagAusstellungLangDEID},
				}).Error; err != nil {
					return err
				}
				if err := tx.Delete(&[]*Tag{
					{ID: TestTagInterviewID},
					{ID: TestTagAusstellungID},
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
	}...)
}
