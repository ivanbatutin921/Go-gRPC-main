package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func (c *DB) Connect() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("Error loading .env file: %s", err.Error())
	// }

	dsn := fmt.Sprintf("postgres://postgres.nmbrrclripkumxfogbfl:P1Ju1C1q7dGOUV1N@aws-0-eu-central-1.pooler.supabase.com:5432/postgres")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
	// 	os.Getenv("PGHOST"),
	// 	os.Getenv("PGUSER"),
	// 	os.Getenv("PGPASSWORD"),
	// 	os.Getenv("PGDATABASE"),
	// 	os.Getenv("PGPORT"),
	// )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	c.db = db
	log.Println("Подключение в бд прошло успешно.")

}
