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

func (p *EchoPlugin) Onstart() {
	fmt.Printf("Starting Plugin %s \n", p.GetId())
}

func (p *EchoPlugin) OnStop() {
	fmt.Printf("Stoping Plugin \n")
}

func (p *EchoPlugin) GetId() string {
	return "echoplugin"
}

func (p *EchoPlugin) Run(message telebot.Message) {
	fmt.Printf("Message recieved for Echo %s \n", message.Text)
	bot := pluginframework.Bot
	if strings.HasPrefix(message.Text, "echo") {
		fmt.Println("Inside Echo")
		bot.SendMessage(message.Chat,
			message.Text, nil)
	}
}
