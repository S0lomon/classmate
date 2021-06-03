package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type interactionHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

// commands will add all the app commands for the bot. Any new commands should go here.
// The parameter addCommand is a function to call to add a command.
func commands(addCommand func(cmd discordgo.ApplicationCommand, h interactionHandler)) {

	addCommand(discordgo.ApplicationCommand{
		Name:        "test",
		Description: "test if the bot is working",
		Options:     nil,
	}, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: "Hello",
			},
		})

		if err != nil {
			log.Println("Error responding to command:", err)
		}
	})
}
