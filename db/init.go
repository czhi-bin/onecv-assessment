package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/czhi-bin/onecv-assessment/config"
	"github.com/czhi-bin/onecv-assessment/model"
	"github.com/czhi-bin/onecv-assessment/utils"
)

var DB *gorm.DB

// Initialize database connection and load database schema
func Init() {
	utils.Logger.Info("Initializing database connection")

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
		utils.Logger.Fatal(err, "Failed to connect to database")
	}

	// Migrate the schema
	loadDatabase()

	utils.Logger.Info("Database connection initialized. Connected to database: ", 
					dbName, ", on host: ", host, ", port: ", port, ", as user: ", user)
}

func loadDatabase() {
	utils.Logger.Info("Loading database schema")

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

	utils.Logger.Info("Database schema loaded")
}
