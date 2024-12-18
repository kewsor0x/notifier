package config

import (
	"goapi/models"
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
}

var JwtKey []byte
var MailSender models.MailSender 

func GetConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading env file: %v", err)
	}

	var config Config
	config.DB_NAME = os.Getenv("DB_NAME")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_PASS = os.Getenv("DB_PASS")
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_PORT = os.Getenv("DB_PORT")

	MailSender.Host = os.Getenv("EMAIL_HOST")
	MailSender.Port = os.Getenv("EMAIL_PORT")
	MailSender.Username = os.Getenv("EMAIL_USERNAME")
	MailSender.Password = os.Getenv("EMAIL_PASSWORD")

	// temporal key to check
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		log.Fatalf("error loading key")
	}

	JwtKey = []byte(jwtKey)

	return config
}
