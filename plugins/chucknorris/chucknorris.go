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

type ChuckNorrisPlugin struct {
	name        string
	command     string
	id          string
	description string
}

func init() {
	pluginframework.Register(&ChuckNorrisPlugin{
		name:        "Chuck Norris Plugin",
		command:     "/chucknorris",
		id:          "[chucknorris]",
		description: "Get random Chuck Norris Jokes. Example /chucknorris ==> gives random chuck norris joke of the day",
	})
}

func (p *ChuckNorrisPlugin) Command() string {
	return p.command
}

func (p *ChuckNorrisPlugin) Description() string {
	return p.description
}

func (p *ChuckNorrisPlugin) OnStart() {
}

func (p *ChuckNorrisPlugin) OnStop() {
}

func (p *ChuckNorrisPlugin) PluginId() string {
	return p.id
}

func (p *ChuckNorrisPlugin) Run(message telebot.Message) {
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
