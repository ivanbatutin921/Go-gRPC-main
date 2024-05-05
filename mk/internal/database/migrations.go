package database

import (
	"root/mk/internal/model"
)

func Migration() {

	if err := DB.DB.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
}
