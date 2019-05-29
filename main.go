package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.Parse()
}

func main() {

	discord, err := discordgo.New("Bot " + Token)

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
