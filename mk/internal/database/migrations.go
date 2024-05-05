package database

import (
	"log"

	"gorm.io/gorm"

	"root/mk/internal/model"
)

func (c *DB) Migration() {
	migration := func(db *gorm.DB) error {
		return db.AutoMigrate(
			&model.User{},
		)
	}

	if err := migration(c.db); err != nil {
		log.Fatal(err)
	}
}
