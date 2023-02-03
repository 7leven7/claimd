package configs

import (
	"fmt"
	"github.com/7leven7/claimd/app/db"
	"github.com/7leven7/claimd/app/models"
)

// PGConnection connecting to a database using the configuration loaded from a local file. It returns an error if the config file cannot be loaded or if the database connection fails.
func PGConnection() error {
	config, err := db.LoadConfig(".")

	if err != nil {
		return fmt.Errorf("Could not load environment variables: %v", err)
	}

	err = db.ConnectDB(&config)
	if err != nil {
		return fmt.Errorf("Could not connect to database: %v", err)
	}
	return nil
}

// Migrate database migrations for the given model
func Migrate() {
	db.DB.AutoMigrate(&models.Instagram{}, &models.Tiktok{}, &models.Platform{})
	fmt.Println("? Migration complete")
}
