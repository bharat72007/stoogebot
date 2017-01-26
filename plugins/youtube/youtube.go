package youtube

import (
	"fmt"
	"github.com/bharat72007/gorest"
	"github.com/stoogebot/pluginframework"
	"github.com/tucnak/telebot"
	"os"
	"strings"
)

const (
	baseapi  = "https://www.googleapis.com"
	watchurl = "https://youtube.com/watch?v="
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

type YoutubePlugin struct{ name string }

func init() {

	//Register Plugin
	fmt.Printf("Initing Youtube Plugin \n")
	pluginframework.Register(&YoutubePlugin{name: "Youtube Plugin"})
}

func (p *YoutubePlugin) OnStart() {
	/*if ApiKey == nil || ApiKey == "" {
		panic("Youtube API Token is not defined.")
	}*/
	fmt.Printf("Starting Youtube %s \n", p.name)
}

func (p *YoutubePlugin) OnStop() {
	fmt.Printf("Stoping Plugin \n")
}

func (p *YoutubePlugin) GetId() string {
	return "YoutubePlugin"
}

func (p *YoutubePlugin) Run(message telebot.Message) {
	ApiKey := os.Getenv("YOUTUBE_KEY")
	fmt.Println(ApiKey)
	fmt.Printf("youtube video to be searched %s \n", message.Text)
	bot := pluginframework.Bot
	if strings.HasPrefix(message.Text, "/youtube") {
		searchKeyword := strings.Replace(strings.Replace(message.Text, "/youtube", "", -1), " ", "", -1)
		rest := gorest.New()
		query := map[string]string{"part": "snippet",
			"type": "video",
			"key":  ApiKey,
			"q":    searchKeyword,
		}
		request, _ := rest.Base(baseapi).Path("youtube").Path("v3").Path("search").Query(query).Get().Request()
		response, _ := rest.Send(request)
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

			bot.SendMessage(message.Chat, title, nil)
			pluginframework.SendVideo(url, message, bot)
			/*			pluginframework.SendPhoto(thumbnailurl, message, nil)
						pluginframework.SendVideo(thumbnailurl, message, nil)
			*/
		} else {
			bot.SendMessage(message.Chat, "Some problem with Youtube", nil)
		}
	}
}
