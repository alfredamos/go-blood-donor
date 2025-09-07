package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	//----> Load environment variables.
	err := godotenv.Load()

	//----> Check for error.
	if err != nil {
		log.Fatal("Error loading env file!")
	}
}
