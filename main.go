package main

import (
	"log"

	"web-scraper/cmd"
	// "web-scraper/modules"

	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
  if err != nil { 
    log.Fatal("Error loading .env file")
  }
	// modules.GetHackathonData()
	cmd.Execute()
}

