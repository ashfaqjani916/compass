package main

import (
	"log"
	"web-scraper/gemini"
	"web-scraper/utils"
"github.com/joho/godotenv"
)

// "web-scraper/cmd"
// "web-scraper/modules"
// "web-scraper/gemini"

func main() {

	err := godotenv.Load()
  if err != nil { 
    log.Fatal("Error loading .env file")
  }
	// cmd.Execute()

	// modules.GetHackathonData()
	utils.Data()
	gemini.SummariseData()
}

// }
