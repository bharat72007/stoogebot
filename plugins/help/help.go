package help

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
	"github.com/tucnak/telebot"
	"strings"
)

type HelpPlugin struct {
	name        string
	command     string
	id          string
	description string
}

func init() {
	pluginframework.Register(&HelpPlugin{
		name:        "HelpPlugin",
		command:     "/help",
		id:          "[help]",
		description: "find out all the plugins available"})
}

func (p *HelpPlugin) OnStart() {
}

func (p *HelpPlugin) OnStop() {
}

func (p *HelpPlugin) PluginId() string {
	return "[help]"
}

func (p *HelpPlugin) Command() string {
	return p.command
}

func (p *HelpPlugin) Description() string {
	return p.description
}

func (p *HelpPlugin) Run(message telebot.Message) {
	bot := pluginframework.Bot
	if strings.HasPrefix(message.Text, p.Command()) {
		var helpstring string
		var plugins map[string]pluginframework.BotPlugin
		plugins = pluginframework.RegisteredPlugins
		for _, plugin := range plugins {
			helpstring = helpstring + plugin.Command() + ": "
			helpstring = helpstring + plugin.Description() + "\n"
		}
		bot.SendMessage(message.Chat, helpstring, nil)
	}
}
