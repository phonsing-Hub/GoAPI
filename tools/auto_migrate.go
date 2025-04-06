package main

import (
	"fmt"
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/phonsing-Hub/GoAPI/internal/models" 
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	dsn := os.Getenv("PGSQL_URI")
	if dsn == "" {
		log.Fatal("PGSQL_URI not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	if err := AutoMigrateAll(db, models.Models); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	
	fmt.Println("Auto migration completed")
}

func AutoMigrateAll(db *gorm.DB, models []interface{}) error {
	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			return err
		}
	}
	return nil
}