// generated from OrmFileSetupTemplate
package orm

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// SetupModels connects to the sqlite database
func SetupModels(logMode bool, filepath string) *gorm.DB {
	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gongleaflet_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filepath), gormConfig)

	if err != nil {
		panic("Failed to connect to database!")
	}

	AutoMigrate(db)

	return db
}

// AutoMigrate migrates db with with orm Struct
func AutoMigrate(db *gorm.DB) {
	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gongleaflet_go_", // table name prefix
	}

	err := db.AutoMigrate( // insertion point for reference to structs
		&DivIconDB{},
		&LayerGroupDB{},
		&MarkerDB{},
		&VisualCircleDB{},
		&VisualLineDB{},
		&VisualMapDB{},
		&VisualTrackDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gongleaflet/go")
	}
	log.Printf("Database Migration of package github.com/fullstack-lang/gongleaflet/go is OK")

	BackRepo.init(db)
}

func ResetDB(db *gorm.DB) { // insertion point for reference to structs
	db.Delete(&DivIconDB{})
	db.Delete(&LayerGroupDB{})
	db.Delete(&MarkerDB{})
	db.Delete(&VisualCircleDB{})
	db.Delete(&VisualLineDB{})
	db.Delete(&VisualMapDB{})
	db.Delete(&VisualTrackDB{})
}