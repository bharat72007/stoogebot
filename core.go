package main

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
	_ "github.com/stoogebot/plugins/chucknorris"
	_ "github.com/stoogebot/plugins/echo"
	"github.com/tucnak/telebot"
	"os"
	"time"
)

func main() {

	fmt.Println("Starting Plugins")
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

/*func HandleQueries(bot *telebot.Bot) {
	for query := range bot.Queries {
		fmt.Println("--- new query ---")
		fmt.Println("from:", query.From.Username)
		fmt.Println("text:", query.Text)

		// Create an article (a link) object to show in results.
		article := &telebot.InlineQueryResultArticle{
			Title: "Yourbot",
			URL:   "https://github.com/tucnak/telebot",
			InputMessageContent: &telebot.InputTextMessageContent{
				Text:           "Telebot is a Telegram bot framework.",
				DisablePreview: false,
			},
		}

		// Build the list of results (make sure to pass pointers!).
		results := []telebot.InlineQueryResult{article}

		// Build a response object to answer the query.
		response := telebot.QueryResponse{
			Results:    results,
			IsPersonal: true,
		}

		// Send it.
		if err := bot.AnswerInlineQuery(&query, &response); err != nil {
			fmt.Println("Failed to respond to query:", err)
		}
	}
}*/
