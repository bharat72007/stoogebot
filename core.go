package main

import (
	"fmt"
	"github.com/stoogebot/pluginframework"
	_ "github.com/stoogebot/plugins/free"
	_ "github.com/stoogebot/plugins/plain"
	"time"
)

func main() {
	fmt.Println("Strating Plugins")
	for _, plugin := range pluginframework.RegisteredPlugins {
		go plugin.Onstart()
	}
	time.Sleep(5000)
}
