package command

import (
	"arb-template/bot"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
)

const (
	// Command names
	Main = "Main"
)

var commonds = []string{Main}

func Run() {
	prompt := promptui.Select{
		Label: "Select Function",
		Items: commonds,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case Main:
		if engine, err := bot.New(); err == nil {
			engine.Run()
		} else {
			log.Printf("Failed to create bot: %v", err)
		}

	default:
		fmt.Println("Unknown")
	}
}
