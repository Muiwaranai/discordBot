package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Run(token string) {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error:", err)
	}

	discord.AddHandler(defaultCommandsHandler)
	discord.AddHandler(adminCommandsHandler)

	discord.Open()
	defer discord.Close()

	fmt.Println("Bot runing")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func defaultCommandsHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID || message.Content[0] != '!' {
		return
	}

	command := strings.ToLower(message.Content[1:])

	switch command {
	case "ping":
		discord.ChannelMessageSendReply(message.ChannelID, "pong", message.Reference())
	}
}

func adminCommandsHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID != "1094657886434635926" || message.Content[0] != '!' {
		return
	}

	command := strings.ToLower(message.Content[1:])

	switch command {
	case "admin":
		discord.ChannelMessageSendReply(message.ChannelID, "hi", message.Reference())
	}
}
