package main

import (
	"log"

	"arb-template/command"
	"arb-template/config"
)

func main() {
	if err := config.Load("config.yml"); err != nil {
		config.GenerateDefaultConfig()
		log.Println(err)
		return
	}

	command.Run()
}
