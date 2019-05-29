package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func main() {

	discord, err := discordgo.New("Bot " + "dummy token")

	if err != nil {
		fmt.Println("Error creating bot,", err)
		return
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()

	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running")

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

}
