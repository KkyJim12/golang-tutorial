package migrations

import (
	"jimmytechnology-golang/config"
	"log"

	"gopkg.in/gormigrate.v1"
)

func Migrate() {
	db := config.GetDB()
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		m1652181450CreateArticlesTable(),
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}
