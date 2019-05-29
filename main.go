package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func main() {

	discord, err := discordgo.New("Bot " + "token")

	fmt.Println(err)
	fmt.Println(discord)
}
