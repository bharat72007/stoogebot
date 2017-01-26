package main

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
	_ "github.com/stoogebot/plugins/chucknorris"
	_ "github.com/stoogebot/plugins/echo"
	_ "github.com/stoogebot/plugins/xkcd"
	_ "github.com/stoogebot/plugins/youtube"
	"github.com/tucnak/telebot"
	"os"
	"time"
)

/*const (
	baseapi  = "https://www.googleapis.com"
	watchurl = "youtube.com/watch?v="
)

type YoutubeResponse struct {
	Kind          string   `json:"-"`
	Etag          string   `json:"-"`
	NextPageToken string   `json:"-"`
	Regioncode    string   `json:"-"`
	PageInfo      pageinfo `json:"-"`
	Items         []item   `json:"items"`
}

type pageinfo struct {
	TotalResults   int `json:"-"`
	ResultsperPage int `json:"-"`
}

type item struct {
	Kind    string  `json:"-"`
	Etag    string  `json:"-"`
	Id      id      `json:"id"`
	Snippet snippet `json:"snippet"`
}

type id struct {
	Kind    string `json:"-"`
	VideoId string `json:"videoId"`
}

type snippet struct {
	PublishedAt          string    `json:"-"`
	ChannelId            string    `json:"-"`
	Title                string    `json:"title"`
	Description          string    `json:"-"`
	ChannelTitle         string    `json:"-"`
	LiveBroadcastContent string    `json:"-"`
	Thumbnails           thumbnail `json:"thumbnails"`
}

type thumbnail struct {
	Default thumb `json:"default"`
	Medium  thumb `json:"medium"`
	High    thumb `json:"high"`
}

type thumb struct {
	Url    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

func main() {
	ApiKey := os.Getenv("YOUTUBE_KEY")
	fmt.Println(ApiKey)
	rest := gorest.New()
	query := map[string]string{"part": "snippet",
		"type": "video",
		"key":  ApiKey,
		"q":    "dangal",
	}
	request, _ := rest.Base(baseapi).Path("youtube").Path("v3").Path("search").Query(query).Get().Request()
	fmt.Println(request)
	response, _ := rest.Send(request)
	fmt.Println(response)
	if response.StatusCode == 200 {
		var youtubeResponse YoutubeResponse
		err := gorest.Response(response, &youtubeResponse, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(youtubeResponse)
		fmt.Println(youtubeResponse.Items)
		fmt.Println(len(youtubeResponse.Items))
		thumbnailurl := youtubeResponse.Items[0].Snippet.Thumbnails.Default.Url
		title := youtubeResponse.Items[0].Snippet.Title
		videoId := youtubeResponse.Items[0].Id.VideoId
		url := watchurl + videoId
		fmt.Println(url)
		fmt.Println(videoId)
		fmt.Println(title)
		fmt.Println(thumbnailurl)
		/*	bot.SendMessage(message.Chat, title, nil)
			bot.SendVideo(url, message, nil)
			bot.SendPhoto(thumbnailurl, message, nil)
			bot.SendVideo(thumbnailurl, message, nil)

	} else {
		fmt.Println("Some Probelm")
		//	bot.SendMessage(message.Chat, "Some problem with Youtube", nil)
	}
}*/

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
