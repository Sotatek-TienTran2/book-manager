package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ConnectPort() string {
	err := godotenv.Load("././.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}
	return ":" + os.Getenv("PORT")
}
