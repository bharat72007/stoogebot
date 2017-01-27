package pluginframework

import (
	"fmt"
	"github.com/tucnak/telebot"
)

type BotPlugin interface {
	OnStart()
	OnStop()
	Run(telebot.Message)
	Command() string
	PluginId() string
	Description() string
}

var RegisteredPlugins = map[string]BotPlugin{}
var Bot *telebot.Bot

func Register(botPlugin BotPlugin) {
	RegisteredPlugins[botPlugin.PluginId()] = botPlugin
}
