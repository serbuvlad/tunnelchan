package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// <1>: Config file
	cfgfile := flag.String("cfg", "", "Config file")
	flag.Parse()

	if *cfgfile == "" {
		fmt.Fprintln(os.Stderr, "no config file provided")
		os.Exit(2)
	}

	cfg, err := parseConfig(*cfgfile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	// <2>: IRC & Discord
	ircData := ircConnect(cfg)
	discordData := discordConnect(cfg)

	for ircChannel, discordChannel := range cfg.Channels {
		ircData.addChannel(ircChannel, discordData.session, discordChannel)
		discordData.addChannel(discordChannel, ircData.conn, ircChannel)
	}

	// <3>: Block - we exit on fatal error or interrupt
	block := make(chan struct{})
	<-block
}
