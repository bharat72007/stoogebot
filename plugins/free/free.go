package free

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
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
