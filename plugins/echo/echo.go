package echo

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
	"github.com/tucnak/telebot"
	"strings"
)

type EchoPlugin struct {
	name        string
	command     string
	id          string
	description string
}

func init() {
	pluginframework.Register(&EchoPlugin{
		name:        "Echo Plugin",
		command:     "/echo",
		id:          "[echo]",
		description: "Echo back the text. Example /echo 'texting to you'",
	})
}

func (p *EchoPlugin) Command() string {
	return p.command
}

func (p *EchoPlugin) Description() string {
	return p.description
}

func (p *EchoPlugin) OnStart() {
}

func (p *EchoPlugin) OnStop() {
}

func (p *EchoPlugin) PluginId() string {
	return p.id
}

func (p *EchoPlugin) Run(message telebot.Message) {
	bot := pluginframework.Bot
	if strings.HasPrefix(message.Text, p.Command()) {
		bot.SendMessage(message.Chat, strings.TrimPrefix(message.Text, p.Command()), nil)
	}
}
