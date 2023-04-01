package bot

import (
	"log"

	"arb-template/internal/types"
)

func (b *Bot) executor(matched *types.MatchedOrders) {
	log.Printf("executor matched orders: %v", matched)

	//processing logic

	//notify
}
