package bot

import (
	"fmt"

	"strings"

	"github.com/JLSELLORSIII/ParakeetGo/config"

	"github.com/bwmarrin/discordgo"
)

var Id string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	Id = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == Id {
		return
	}

	// if m.content contains botid (Mentions) and "ping" then send "pong!"
	if m.Content == "<@"+Id+"> ping" || m.Content == "<@"+Id+"> ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}

	if m.Content == config.BotPrefix+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}

	if strings.Contains(m.Content, config.BotPrefix+"create-corpus") {
		CreateCorpus(s, m)
	}

}
