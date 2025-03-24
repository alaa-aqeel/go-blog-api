package database

import (
	"fmt"
	"log"
	"os"

	"github.com/alaa-aqeel/govalid/src/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to database successfully")
	DB = db
}

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migration completed successfully")
}

func RollbackDatabase(db *gorm.DB) {
	err := db.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatal("Failed to drop table:", err)
	}

	log.Println("Database migration completed successfully")
}
