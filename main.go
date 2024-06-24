package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	test, exists := os.LookupEnv("TEST")
	if exists {
		fmt.Printf("Found TEST variable %s\n", test)
	} else {
		fmt.Println("No TEST variable found")
	}
}
