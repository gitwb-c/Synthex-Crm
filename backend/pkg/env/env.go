package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() error {
	err := godotenv.Load()
	if err != nil {
		if os.IsNotExist(err) {
			log.Println(".env file not found, using existing environment variables")
			return nil
		}
		log.Printf("error loading .env file: %v", err)
		return err
	}
	log.Println(".env loaded successfully")
	return nil
}
