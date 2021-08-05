package migrations

import (
	"github.com/marqueswsm/web-api-with-golang/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
