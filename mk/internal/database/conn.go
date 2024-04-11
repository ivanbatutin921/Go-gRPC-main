package initializers

import (
	"fmt"
	//"log"
	//"os"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	//"github.com/joho/godotenv"
)

type Data struct {
	DB *gorm.DB
}

var DB Data

func ConnectToDB() (Data, error) {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
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
		//return nil, err
	}
	DB = Data{
		DB: db,
	}
	fmt.Println("connected")
	return DB, nil
}