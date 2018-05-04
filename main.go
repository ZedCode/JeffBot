package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token    string
	pubgTime string
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
		// ignore messages from self
		return
	}
	// Store all commands here...
	// Break up the message so that we are only seeing if it starts with

	fmt.Println(m.Content)
	fmt.Println(len(pubgTime))
	switch dCmd := strings.Split(m.Content, " ")[0]; dCmd {
	case "!pubg":
		if len(pubgTime) > 0 {
			s.ChannelMessageSend(m.ChannelID, pubgTime)
		} else {
			s.ChannelMessageSend(m.ChannelID, "No game scheduled...")
		}
	case "!schedule":
		switch gName := strings.Split(m.Content, " ")[1]; gName {
		case "pubg":
			// Now we have to strip out the first two parts of the string:
			pubgTime = strings.Replace(m.Content, "!schedule pubg ", "", -1)
		}
	}
}
