package main

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
	_ "github.com/stoogebot/plugins/free"
	_ "github.com/stoogebot/plugins/plain"
	"github.com/tucnak/telebot"
	"os"
	//"strings"
	"time"
)

func main() {
	fmt.Println("Starting Plugins")
	for _, plugin := range pluginframework.RegisteredPlugins {
		go plugin.Onstart()
	}
	//time.Sleep(5000)

	//Environment Variables to GET values for registered Chat Bot.
	//telebotname := os.Getenv("TELEBOT_NAME")
	telebottoken := os.Getenv("TELEBOT_TOKEN")
	bot, err := telebot.NewBot(telebottoken)

	if err != nil {
		fmt.Print(err)
		return
	}

	messages := make(chan telebot.Message, 100)
	bot.Listen(messages, 1*time.Second)
	pluginframework.Bot = bot

	for message := range messages {
		for _, plugin := range pluginframework.RegisteredPlugins {
			go plugin.Run(message)
		}
	}
}
