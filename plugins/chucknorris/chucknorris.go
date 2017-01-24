package chucknorris

import (
	"fmt"
	"github.com/bharat72007/gorest"
	"github.com/stoogebot/pluginframework"
	"github.com/tucnak/telebot"
	"strings"
)

const (
	baseapi = "https://api.chucknorris.io/jokes/"
)

type Chucknorriscall struct {
	Category string `json:"category"`
	Icon_url string `json:"icon_url"`
	Id       string `json:"id"`
	Url      string `json:"url"`
	Value    string `json:"value"`
}

type ChuckNorrisPlugin struct{ name string }

func init() {
	//Register Plugin
	fmt.Printf("Inviting ChuckNorrisPlugin \n")
	pluginframework.Register(&ChuckNorrisPlugin{name: "One and Only one ChuckNorris"})
}

func (p *ChuckNorrisPlugin) OnStart() {
	fmt.Printf("Starting ChuckNorris %s \n", p.name)
}

func (p *ChuckNorrisPlugin) OnStop() {
	fmt.Printf("Stoping Plugin \n")
}

func (p *ChuckNorrisPlugin) GetId() string {
	return "Plugins Creator ChuckNorris"
}

func (p *ChuckNorrisPlugin) Run(message telebot.Message) {
	fmt.Printf("Message recieved for Echo %s \n", message.Text)
	bot := pluginframework.Bot
	if strings.HasPrefix(message.Text, "/chucknorris") {
		rest := gorest.New()
		request, _ := rest.Base(baseapi).Path("random").Get().Request()
		response, _ := rest.Send(request)
		if response.StatusCode == 200 {
			var resp Chucknorriscall
			gorest.Response(response, &resp, nil)
			chuckResponse := resp.Value
			bot.SendMessage(message.Chat, chuckResponse, nil)
		} else {
			bot.SendMessage(message.Chat, "Error: Server not working, Chuck Norris took Over", nil)
		}
	}
}
