package plain

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
)

type PlainPlugin struct{}

func init() {
	fmt.Printf("Registering Plain \n")
	pluginframework.Register(&PlainPlugin{})
}

func (p *PlainPlugin) Onstart() {
	fmt.Printf("Starting Plugin %s \n", p.GetId())
}

func (p *PlainPlugin) OnStop() {
	fmt.Printf("Stoping Plugin \n")
}

func (p *PlainPlugin) GetId() string {
	return "plainplugin"
}