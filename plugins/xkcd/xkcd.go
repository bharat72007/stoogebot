package xkcd

import (
	"fmt"
	"github.com/bharat72007/gorest"
	"github.com/stoogebot/pluginframework"
	"github.com/tucnak/telebot"
	"strings"
)

const (
	baseapi = "https://xkcd.com/"
)

type Xkcd struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Safe_title string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:day"`
}

type XkcdPlugin struct{ name string }

func init() {
	//Register Plugin
	fmt.Printf("Starting XkcdPlugin \n")
	pluginframework.Register(&XkcdPlugin{name: "XKCD"})
}

func (p *XkcdPlugin) OnStart() {
	fmt.Printf("Starting Xkcd %s \n", p.name)
}

func (p *XkcdPlugin) OnStop() {
	fmt.Printf("Stoping Plugin \n")
}

func (p *XkcdPlugin) GetId() string {
	return "XcdPlugin"
}

func (p *XkcdPlugin) Run(message telebot.Message) {
	fmt.Println("Message xkcd" + message.Text)
	bot := pluginframework.Bot
	if strings.HasPrefix(message.Text, "/xkcd") {
		rest := gorest.New()
		//Trim all spaces, after trimming "/xkcd" from message test.
		//Example /xkcd 22 ==> 22
		number := strings.Replace(strings.Replace(message.Text, "/xkcd", "", -1), " ", "", -1)
		fmt.Println(number)
		request, _ := rest.Base(baseapi).Path(number).Path("info.0.json").Get().Request()
		fmt.Println(request)
		response, _ := rest.Send(request)
		fmt.Println(response)
		if response.StatusCode == 200 {
			var xkcd Xkcd
			gorest.Response(response, &xkcd, nil)
			//Xkcd Image URL.
			xkcdImgUrl := xkcd.Img
			fmt.Println(xkcdImgUrl)
			//bot.SendMessage(message.Chat, xkcdImgUrl, nil)
			pluginframework.SendPhoto(xkcdImgUrl, message, bot)
			pluginframework.SendVideo("https://www.youtube.com/watch?v=f6kdp27TYZs", message, bot)
			pluginframework.SendAudio("https://github.com/eternnoir/gotelebot/blob/master/test_data/record.mp3", message, bot)
			pluginframework.SendSticker("https://github.com/eternnoir/gotelebot/blob/master/test_data/go.webp", message, bot)

		} else {
			bot.SendMessage(message.Chat, "Some problem with Xkcd", nil)
		}
	}
}
