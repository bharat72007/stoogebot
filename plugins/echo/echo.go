package echo

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
	"github.com/tucnak/telebot"
	"strings"
)

type EchoPlugin struct{}

func init() {
	fmt.Printf("Registering Echo Plugin \n")
	pluginframework.Register(&EchoPlugin{})
}

func (p *EchoPlugin) OnStart() {
	fmt.Printf("Starting Plugin %s \n", p.GetId())
}

func (p *EchoPlugin) OnStop() {
	fmt.Printf("Stoping Plugin \n")
}

func (p *EchoPlugin) GetId() string {
	return "echoplugin"
}

func (p *EchoPlugin) Run(message telebot.Message) {
	bot := pluginframework.Bot
	if strings.HasPrefix(message.Text, "/echo") {
		bot.SendMessage(message.Chat, strings.TrimPrefix(message.Text, "/echo"), nil)
	}
}
