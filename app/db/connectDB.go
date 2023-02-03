package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB(config *Config) error {
	var err error
	//err := godotenv.Load("./.env")
	//if err != nil {
	//	log.Fatalf("Some error occured. Err: %s", err)
	//}
	//
	//host := os.Getenv("DB_HOST")
	//user := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASSWORD")
	//dbname := os.Getenv("DB_NAME")
	//port := os.Getenv("DB_PORT")
	//
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	dsn := fmt.Sprintf("host=postgres user=claimd password=claimdbko dbname=claimd port=5432 sslmode=disable")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")

	return nil
}
