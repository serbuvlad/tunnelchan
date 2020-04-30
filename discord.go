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
	data := &discordData{}
	data.channels = make(map[string]chan *discord.MessageCreate)

	var err error
	data.session, err = discord.New(fmt.Sprintf("Bot %s", cfg.Discord.Token))
	if err != nil {
		log.Fatal(err)
	}

	err = data.session.Open()
	if err != nil {
		log.Fatal(err)
	}

	data.session.AddHandler(func(session *discord.Session, message *discord.MessageCreate) {
		ch, ok := data.channels[message.ChannelID]
		if ok {
			ch <- message
		}
	})

	me, err := data.session.User("@me")
	if err != nil {
		log.Fatal(err)
	}
	data.whoami = me.ID

	return data
}

func (data *discordData) addChannel(discordChannel string, ircConn *irc.Conn, ircChannel string) {
	ch := make(chan *discord.MessageCreate)
	data.channels[discordChannel] = ch

	go func() {
		for message := range ch {
			if message.Author.ID != data.whoami {
				ircConn.Privmsgf(ircChannel, "<%s> %s", message.Author.Username, message.Content)
			}
		}
	}()
}
