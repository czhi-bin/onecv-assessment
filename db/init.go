package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/czhi-bin/onecv-assessment/config"
	"github.com/czhi-bin/onecv-assessment/model"
)

var DB *gorm.DB

// Initialize database connection and load database schema
func Init() {
	host := config.PSQL_HOST
	user := config.PSQL_USER
	password := config.PSQL_PASSWORD
	port := config.PSQL_PORT
	dbName := config.PSQL_DBNAME

	var err error
	DB, err = gorm.Open(
		postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable timezone=Asia/Singapore",
			host, user, password, dbName, port,
		)),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err, "Failed to connect to database")
	}

	// Migrate the schema
	loadDatabase()
}

func loadDatabase() {
	err := DB.AutoMigrate(&model.Teacher{})
	if err != nil {
		log.Fatal(err, "Failed to migrate teacher table")
	}

	err = DB.AutoMigrate(&model.Student{})
	if err != nil {
		log.Fatal(err, "Failed to migrate student table")
	}

	err = DB.AutoMigrate(&model.Registration{})
	if err != nil {
		log.Fatal(err, "Failed to migrate registration table")
	}

}
