package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token string
)

func init() {
	flag.StringVar(&token, "t", "", "Token")
}

func main() {
	flag.Parse()
	jeffBot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Failed to start bot: ", err.Error())
	}
	jeffBot.AddHandler(messageHandler)
	err = jeffBot.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	jeffBot.Close()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Contains(m.Content, "!jeff") {
		s.ChannelMessageSend(m.ChannelID, "Who's the sexiest man alive?")
	}
}
