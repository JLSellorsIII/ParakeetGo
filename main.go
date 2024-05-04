package main

import (
	"fmt"
	"github.com/JLSELLORSIII/ParakeetGo/bot"
	"github.com/JLSELLORSIII/ParakeetGo/config"
)

// Reads config and starts the bot
func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Start()

	<-make(chan struct{})

	return
}
