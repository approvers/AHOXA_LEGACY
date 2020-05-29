package main

import (
	"AHOXA/src"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	discord, Err := discordgo.New()
	discord.Token = BOT_TOKEN
	if Err != nil {
		fmt.Println("Error log in")
		fmt.Println(Err)
	}

	discord.AddHandler(src.MessageCreate)

	Err = discord.Open()
	if Err != nil {
		fmt.Println(Err)
	}

	fmt.Println("Ready Bot.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- sc
	return
}
