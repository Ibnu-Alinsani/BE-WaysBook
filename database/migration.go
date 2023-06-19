package database

import (
	"waysbook/models"
	"waysbook/pkg/postgresql"
)

func RunMigration() {
	err := postgresql.DB.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Cart{},
		&models.Transaction{},
	)

	if err != nil {
		panic("Migration Failed")
	}

}
