package migrations

import (
	"jimmytechnology-golang/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1652181450CreateArticlesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1652181450",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Article{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("articles").Error
		},
	}
}
