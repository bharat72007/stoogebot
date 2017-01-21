package free

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
	"github.com/tucnak/telebot"
	"strings"
)

type FreePlugin struct{ name string }

func init() {
	//Register Plugin
	fmt.Printf("Registering Free \n")
	pluginframework.Register(&FreePlugin{name: "Free"})
}

func (p *FreePlugin) Onstart() {
	fmt.Printf("Starting Plugin %s \n", p.name)
}

func (p *FreePlugin) OnStop() {
	fmt.Printf("Stoping Plugin \n")
}

func (p *FreePlugin) GetId() string {
	return "freeplugin"
}

func (p *FreePlugin) Run(message telebot.Message) {
	fmt.Printf("Message recieved for Echo %s \n", message.Text)
	bot := pluginframework.Bot
	if strings.HasPrefix(message.Text, "echo") {
		fmt.Println("Inside Echo")
		bot.SendMessage(message.Chat,
			message.Text, nil)
	}
}
