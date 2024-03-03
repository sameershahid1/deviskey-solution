package database

import (
	"database/sql"
	"dev-solution/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB
var SqlDB *sql.DB

func ConnectToDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	SqlDB, err := GormDB.DB()
	if err != nil {
		log.Fatal("Database crash", err)
	}

	SqlDB.SetMaxOpenConns(10)
	SqlDB.SetMaxIdleConns(5)

	fmt.Println("Database Connected")
}

func MigrateData() {
	err := GormDB.AutoMigrate(&model.VehiclePart{})
	if err != nil {
		log.Fatal("Failed to migrate", err)
	}

	fmt.Println("Data is Migrated")
}
