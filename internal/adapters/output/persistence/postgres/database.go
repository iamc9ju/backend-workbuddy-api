package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDatagase() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	//database creddentails from .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", dbHost, dbPort, dbName, dbUser, dbPassword)

	env := os.Getenv("APP_ENV")

	var logLevel logger.LogLevel
	switch env {

	case "development":
		logLevel = logger.Info //dev mode show sql
	case "stagign":
		logLevel = logger.Warn //less logs
	default:
		logLevel = logger.Error //prod mode -> show warning , error
	}

	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //show in terminal,/ln every line ,show with timestamps
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold in terminal
			LogLevel:                  logLevel,    // Log level
			Colorful:                  false,       // Disable color
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: sqlLogger,
		DryRun: false,
	})
	if err != nil {
		return nil, err
	}

	return db, nil

}
