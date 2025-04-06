package main

import (
	"log"
	"os"

	"discordBot/bot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../build/.env")
	if err != nil {
		log.Println("Error with setting env", err)
	}

	bot.Run(os.Getenv("token"))
}
