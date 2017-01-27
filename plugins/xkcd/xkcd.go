package xkcd

import (
	"fmt"
	"github.com/bharat72007/gorest"
	"github.com/stoogebot/pluginframework"
	"github.com/tucnak/telebot"
	"strconv"
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

type XkcdPlugin struct {
	name        string
	command     string
	id          string
	description string
}

func init() {
	pluginframework.Register(&XkcdPlugin{
		name:        "XKCD",
		command:     "/xkcd",
		id:          "[xkcd]",
		description: "Get xkcd comic Image based on Id. Example /xkcd 888, Provide comic Image with Id(888) to user",
	})
}

func (p *XkcdPlugin) Command() string {
	return p.command
}

func (p *XkcdPlugin) Description() string {
	return p.description
}

func (p *XkcdPlugin) OnStart() {
}

func (p *XkcdPlugin) OnStop() {
}

func (p *XkcdPlugin) PluginId() string {
	return p.id
}

func (p *XkcdPlugin) Run(message telebot.Message) {
	bot := pluginframework.Bot
	if strings.HasPrefix(message.Text, p.Command()) {
		rest := gorest.New()
		//Trim all spaces, after trimming "/xkcd" from message test.
		//Example /xkcd 22 ==> 22
		number := strings.Replace(strings.Replace(message.Text, p.Command(), "", -1), " ", "", -1)
		//Convert the String to Integer, Using strconv and check whether valid number or not.
		num, err := strconv.Atoi(number)
		if err != nil {
			bot.SendMessage(message.Chat, "Error : Id is not an Integer", options)
		} else if num < 100 && num > 1100 {
			bot.SendMessage(message.Chat, "Error Id should be between [100,1100]", options)
		} else {
			request, _ := rest.Base(baseapi).Path(num).Path("info.0.json").Get().Request()
			response, _ := rest.Send(request)
			if response != nil; response.StatusCode == 200 {
				var xkcd Xkcd
				gorest.Response(response, &xkcd, nil)
				//Xkcd Image URL.
				xkcdImgUrl := xkcd.Img
				fmt.Println(xkcdImgUrl)
				pluginframework.SendPhoto(xkcdImgUrl, message, bot)
				/*			pluginframework.SendVideo("https://www.youtube.com/watch?v=f6kdp27TYZs", message, bot)
							pluginframework.SendAudio("https://github.com/eternnoir/gotelebot/blob/master/test_data/record.mp3", message, bot)
							pluginframework.SendSticker("https://github.com/eternnoir/gotelebot/blob/master/test_data/go.webp", message, bot)
				*/
			} else {
				bot.SendMessage(message.Chat, "Some problem with Xkcd, Try After Sometime", nil)
			}
		}
	}
}
