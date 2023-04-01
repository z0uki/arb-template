package bot

import (
	"arb-template/internal/sdk/blur"
	"arb-template/internal/sdk/opensea"
	"log"

	"arb-template/config"
	"arb-template/internal/helper"
	"arb-template/internal/types"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Bot struct {
	botWallet     *types.Wallet
	signingWallet *types.Wallet
	ethClient     *ethclient.Client
	blurConn      *blur.Client
	blurWss       *blur.ClientWs
	openseaConn   *opensea.Client
}

func New() (*Bot, error) {
	botWallet, err := helper.WalletFromPrivateKey(config.Conf.PrivateKey)
	if err != nil {
		return nil, err
	}

	signingWallet, err := helper.WalletFromPrivateKey(config.Conf.SigningKey)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(config.Conf.EthNode)
	if err != nil {
		return nil, err
	}

	blurConn, err := blur.New()
	if err != nil {
		return nil, err
	}

	blurWss, err := blur.NewWs()
	if err != nil {
		return nil, err
	}

	openseaConn, err := opensea.New()
	if err != nil {
		return nil, err
	}

	return &Bot{
		botWallet:     botWallet,
		signingWallet: signingWallet,
		ethClient:     ethClient,
		blurConn:      blurConn,
		blurWss:       blurWss,
		openseaConn:   openseaConn,
	}, nil
}

func (b *Bot) Run() {
	log.Printf("Bot started with address: %s", b.botWallet.Address)

	receiveChannel := make(chan *types.MatchedOrders)
	go b.alertOrder()

	for {
		matched := <-receiveChannel

		//async or sync
		go b.executor(matched)
	}
}
