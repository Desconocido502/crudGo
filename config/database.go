package config

import (
	"github.com/Desconocido502/test/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var Database *gorm.DB
var DATABASE_URI string

func Connect() error {
	loadEnvVariables()

	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Dog{})

	return nil
}

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBUser := os.Getenv("DB_USER")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	// Si DBPassword está vacío, no incluirlo en la cadena de conexión
	DBPassword := os.Getenv("DB_PASSWORD")
	if DBPassword != "" {
		DBPassword = ":" + DBPassword
	}

	DATABASE_URI = DBUser + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
