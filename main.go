package main

import (
	"log"

	"github.com/joho/godotenv"
	"oxion.xyz/gomusic/cmd"
)

func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	cmd.Execute()
}
