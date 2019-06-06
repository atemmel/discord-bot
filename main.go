package main

import (
	"os"
	"os/signal"
	"syscall"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
)

var (
	Token string
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.Parse()
	rand.Seed(time.Now().Unix() )
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

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "russian roulette" {
		members, err := s.GuildMembers(m.GuildID, "", 90)

		if( err != nil) {
			return
		}

		member := members[rand.Intn(len(members) ) ]

		length := 3 + rand.Intn(32 - 3)

		name := RandStringRunes(length)

		fmt.Println(member.User, "will have their name changed to", name)

		err = s.GuildMemberNickname(m.GuildID, member.User.ID, name)

		if( err != nil) {
			fmt.Println(err)
		}
	}
}
