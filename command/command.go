package command

import (
	"arb-template/config"
	"arb-template/internal/helper"
	"fmt"
	"log"

	"arb-template/bot"

	"github.com/manifoldco/promptui"
)

const (
	// Command names
	Main    = "Main"
	Approve = "Approve"
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
	case Approve:
		wallet, err := helper.WalletFromPrivateKey(config.Conf.PrivateKey)
		if err != nil {
			log.Printf("Failed to create wallet: %v", err)
		}
		err = approveWETH(wallet)
		if err != nil {
			log.Printf("Failed to approve WETH: %v", err)
		}
		log.Printf("Successfully approved!")
	default:
		fmt.Println("Unknown")
	}
}
