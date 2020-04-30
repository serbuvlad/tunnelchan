package main

import (
	"fmt"
	"log"

	discord "github.com/bwmarrin/discordgo"
	irc "github.com/fluffle/goirc/client"
)

type ircData struct {
	conn *irc.Conn

	channels map[string]chan *irc.Line
}

func ircConnect(cfg *config) *ircData {
	data := &ircData{}
	data.channels = make(map[string]chan *irc.Line)

	data.conn = irc.SimpleClient(cfg.Irc.Nick)

	connected := make(chan struct{})
	data.conn.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) {
		close(connected)
	})

	data.conn.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) {
		log.Fatal("Disconnected from IRC")
	})

	data.conn.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
		chans := line.Args[0]

		// not correct: TODO fix
		{
			ch, ok := data.channels[chans]
			if ok {
				ch <- line
			}
		}
	})

	err := data.conn.ConnectTo(cfg.Irc.Server)
	if err != nil {
		log.Fatal(err)
	}
	<-connected

	return data
}

func (data *ircData) addChannel(ircChannel string, discordSession *discord.Session, discordChannel string) {
	ch := make(chan *irc.Line)
	data.channels[ircChannel] = ch
	data.conn.Join(ircChannel)

	go func() {
		for line := range ch {
			discordSession.ChannelMessageSend(discordChannel, fmt.Sprintf("<%s> %s", line.Nick, line.Text()))
		}
	}()
}
