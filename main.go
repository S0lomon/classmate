package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

func main() {
	token, err := os.ReadFile("bot.token")
	if err != nil {
		log.Fatalln("Failed to read token file:", err)
	}

	session, err := discordgo.New("Bot " + string(token))
	if err != nil {
		log.Fatalln("Failed to establish the bot session:", err)
	}

	err = session.Open()
	if err != nil {
		log.Fatalln("Failed to open the websocket connection to discord:", err)
	}

	log.Println("Starting session as", session.State.User.Username)
}
