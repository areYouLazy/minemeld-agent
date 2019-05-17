package main

import (
	"github.com/areYouLazy/minemeld-agent/agent"
	"github.com/areYouLazy/minemeld-agent/log"
)

func init() {
	log.Init()

	//parse options to get flags from cli
	agent.ParseOptions()

	//start webserver
	go agent.WebServerInit()
}

func main() {
	//start infinite loop
	for {
		//start routine
		agent.FetchEndpoints()
	}
}
