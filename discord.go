package main

import (
	"fmt"
	"log"

	discord "github.com/bwmarrin/discordgo"
	irc "github.com/fluffle/goirc/client"
)

type discordData struct {
	session *discord.Session
	whoami  string

	channels map[string]chan *discord.MessageCreate
}

func discordConnect(cfg *config) *discordData {
	discordData := &discordData{}
	discordData.channels = make(map[string]chan *discord.MessageCreate)

	var err error
	discordData.session, err = discord.New(fmt.Sprintf("Bot %s", cfg.Discord.Token))
	if err != nil {
		log.Fatal(err)
	}

	err = discordData.session.Open()
	if err != nil {
		log.Fatal(err)
	}

	discordData.session.AddHandler(func(session *discord.Session, message *discord.MessageCreate) {
		ch, ok := discordData.channels[message.ChannelID]
		if ok {
			ch <- message
		}
	})

	me, err := discordData.session.User("@me")
	if err != nil {
		log.Fatal(err)
	}
	discordData.whoami = me.ID

	return discordData
}

func (discordData *discordData) addChannel(discordChannel string, ircConn *irc.Conn, ircChannel string) {
	ch := make(chan *discord.MessageCreate)
	discordData.channels[discordChannel] = ch

	go func() {
		for message := range ch {
			if message.Author.ID != discordData.whoami {
				ircConn.Privmsgf(ircChannel, "<%s> %s", message.Author.Username, message.Content)
			}
		}
	}()
}
