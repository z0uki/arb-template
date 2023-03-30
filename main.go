package main

import (
	"arb-template/command"
	"arb-template/config"
	"log"
)

func main() {
	if err := config.Load("config.yml"); err != nil {
		config.GenerateDefaultConfig()
		log.Println(err)
		return
	}

	command.Run()
}
