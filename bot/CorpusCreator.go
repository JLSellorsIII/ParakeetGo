package bot

import (
	"bufio"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
)

const BufferSize int = 4096
const CommandPrefix uint8 = '!'
const StopToken string = "<STOP TOKEN>"

// CreateCorpus Crawls though all the messages in all the channels in all the guilds the bot is in
func CreateCorpus(s *discordgo.Session, m *discordgo.MessageCreate) {
	info := strings.Split(m.Content, " ")
	targetID := info[1]
	// Loop through each guild in the
	file, err := os.Create("corpora/text/" + targetID + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		log.Println("closing corpus file")
		err := file.Close()
		if err != nil {

		}
	}(file)
	writer := bufio.NewWriterSize(file, BufferSize)
	for _, guild := range s.State.Guilds {
		// Get channels for this guild
		channels, _ := s.GuildChannels(guild.ID)
		for _, channel := range channels {
			// Check if channel is a voice channel
			if channel.Type != discordgo.ChannelTypeGuildVoice {
				println("text channel")
				fetchMessages(s, channel.ID, channel.LastMessageID, targetID, writer)
			}
		}
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}

func fetchMessages(s *discordgo.Session, channelID string, lastMessageID string, targetID string, writer *bufio.Writer) {
	firstMessage, err := s.ChannelMessage(channelID, lastMessageID)
	if err != nil {
		log.Println("Error fetching messages")
	} else {
		if firstMessage.Author.ID == targetID && (string(firstMessage.Content)[0] != CommandPrefix) {
			bufferedWrite(writer, firstMessage.Content+StopToken+"\n")
		}
	}
	recursiveFetchMessages(s, channelID, lastMessageID, targetID, writer)
}

// RecursiveFetchMessages recursively fetches and processes messages in a channel
func recursiveFetchMessages(s *discordgo.Session, channelID string, lastMessageID string, targetID string, writer *bufio.Writer) {
	messages, err := s.ChannelMessages(channelID, 100, lastMessageID, "", "")
	if err != nil {
		log.Println("Error fetching messages")
	} else {
		for _, message := range messages {
			if message.Author.ID == targetID && (string(message.Content)[0] != '!') {
				bufferedWrite(writer, message.Content+StopToken+"\n")
			}
		}
		if len(messages) == 100 {
			recursiveFetchMessages(s, channelID, messages[99].ID, targetID, writer)
		}
	}

}

// given
func bufferedWrite(writer *bufio.Writer, content string) {
	contentBytes := []byte(content)
	if len(contentBytes) > writer.Available() {
		err := writer.Flush()
		if err != nil {
			return
		}
	}
	_, err := writer.Write(contentBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
}
