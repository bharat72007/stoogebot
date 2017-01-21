package pluginframework

import (
	"fmt"
	"github.com/tucnak/telebot"
)

type BotPlugin interface {
	Onstart()
	OnStop()
	GetId() string
	Run(telebot.Message)
}

var RegisteredPlugins = map[string]BotPlugin{}
var Bot *telebot.Bot

func Register(botPlugin BotPlugin) {
	RegisteredPlugins[botPlugin.GetId()] = botPlugin
	fmt.Println(RegisteredPlugins)
}
