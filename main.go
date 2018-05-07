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
	"github.com/peterbourgon/diskv"
)

var (
	token     string // Discord Token
	pubgToken string // API Token
	d         *diskv.Diskv
)

func init() {
	flag.StringVar(&token, "t", "", "Discord Token")
	flag.StringVar(&pubgToken, "p", "", "PUBG API Token")
}

func main() {
	flag.Parse()
	// Before we setup the bot, we need to setup the key,value store on disk
	// so that our data isn't simply in memory and lost at each restart of the bot.
	d = diskv.New(diskv.Options{
		BasePath:     "./", // not sure this is a good idea...
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 1024 * 1024,
	})

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
	switch dCmd := strings.Split(m.Content, " ")[0]; dCmd {
	case "!pubg":
		// Try to read the key...
		v, e := d.Read("pubgTime")
		if e != nil {
			s.ChannelMessageSend(m.ChannelID, "No game scheduled...")
		} else {
			s.ChannelMessageSend(m.ChannelID, string(v))
		}
	case "!schedule":
		if len(strings.Split(m.Content, " ")) < 3 {
			s.ChannelMessageSend(m.ChannelID, "You must provide the game you want to schedule. Example !schedule pubg 2pm PST")
			return
		}
		switch gName := strings.Split(m.Content, " ")[1]; gName {
		case "pubg":
			// Now we have to strip out the first two parts of the string:
			if err := d.Write("pubgTime", []byte(strings.Replace(m.Content, "!schedule pubg ", "", -1))); err != nil {
				s.ChannelMessageSend(m.ChannelID, "Failed to schedule a time for some reason. Sorry!")
			} else {
				s.ChannelMessageSend(m.ChannelID, "Game Set! Unset with !del pubg")
			}
		}
	case "!register":
		userData := getPlayerData(strings.Split(m.Content, " ")[1])
		fmt.Println(userData)
		if err := d.Write(m.Author.ID, []byte(userData.Data[0].ID)); err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error Associating your discord user ID with PUBG Username "+strings.Split(m.Content, " ")[1])
			log.Println("Error associating pubg username for user: " + err.Error())
		} else {
			s.ChannelMessageSend(m.ChannelID, "Associated your discord user ID with PUBG Username "+strings.Split(m.Content, " ")[1])
		}
	}
}
