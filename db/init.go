package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/czhi-bin/onecv-assessment/model"
)

var DB *gorm.DB

// Initialize database connection and load database schema
func Init() {
	loadEnv()
	host := os.Getenv("PSQL_HOST")
	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASSWORD")
	port := os.Getenv("PSQL_PORT")
	dbName := os.Getenv("PSQL_DBNAME")

	var err error
	DB, err = gorm.Open(
		postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timezone=Asia/Singapore",
			host, user, password, dbName, port,
		)),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	// Migrate the schema
	loadDatabase()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	DB.AutoMigrate(&model.Teacher{})
	DB.AutoMigrate(&model.Student{})
	DB.AutoMigrate(&model.Registration{})
}
