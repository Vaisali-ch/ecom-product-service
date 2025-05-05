// 2025

package config

import (
	"log"

	"github.com/vaisali-ch/ecom-product-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupProductDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=productdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&models.Product{})
	DB = db
}
