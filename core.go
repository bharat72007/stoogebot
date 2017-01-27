package main

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
	_ "github.com/stoogebot/plugins/chucknorris"
	_ "github.com/stoogebot/plugins/echo"
	_ "github.com/stoogebot/plugins/help"
	_ "github.com/stoogebot/plugins/xkcd"
	_ "github.com/stoogebot/plugins/youtube"
	"github.com/tucnak/telebot"
	"os"
	"time"
)

func main() {

	for _, plugin := range pluginframework.RegisteredPlugins {
		go plugin.OnStart()
	}

	//Environment Variables to GET values for registered Chat Bot.
	//telebotname := os.Getenv("TELEBOT_NAME")
	telebottoken := os.Getenv("TELEBOT_TOKEN")
	fmt.Println(telebottoken)
	bot, err := telebot.NewBot(telebottoken)

	if err != nil {
		fmt.Print(err)
		return
	}

	bot.Messages = make(chan telebot.Message, 100)
	//	bot.Queries = make(chan telebot.Query, 1000)
	pluginframework.Bot = bot
	go HandleMessages(bot)
	//	go HandleQueries(bot)

	//In order to Perform Polling i.e Long Polling Time 1 second
	bot.Start(1 * time.Second)

}

func HandleMessages(bot *telebot.Bot) {
	for message := range bot.Messages {
		for _, plugin := range pluginframework.RegisteredPlugins {
			go plugin.Run(message)
		}
	}
}
