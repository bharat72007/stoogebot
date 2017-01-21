package pluginframework

import (
	"fmt"
)

type BotPlugin interface {
	Onstart()
	OnStop()
	GetId() string
}

var RegisteredPlugins = map[string]BotPlugin{}

func Register(botPlugin BotPlugin) {
	RegisteredPlugins[botPlugin.GetId()] = botPlugin
	fmt.Println(RegisteredPlugins)
}
