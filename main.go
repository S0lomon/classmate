package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
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

	defer session.Close()

	err = session.Open()
	if err != nil {
		log.Fatalln("Failed to open the websocket connection to discord:", err)
	}

	log.Println("Starting session as", session.State.User.Username)

	var handlers = map[string]interactionHandler{}

	// Register all commands
	commands(func(cmd discordgo.ApplicationCommand, h interactionHandler) {
		name := cmd.Name

		_, err = session.ApplicationCommandCreate(session.State.User.ID, "", &cmd)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", name, err)
		}

		_, ok := handlers[name]
		if ok {
			log.Fatalln("Duplicate command registered:", name)
		}

		handlers[name] = h
	})

	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.Data.Name]; ok {
			h(s, i)
		}
	})

	log.Println("Waiting for exit signal")

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
