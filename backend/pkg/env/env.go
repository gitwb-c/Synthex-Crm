package env

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
		return err
	}
	log.Print(".env loaded successfully")
	return nil
}
