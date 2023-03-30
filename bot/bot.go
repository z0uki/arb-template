package bot

import (
	"arb-template/config"
	"arb-template/internal/helper"
	"arb-template/internal/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Bot struct {
	wallet *types.Wallet

	ethClient *ethclient.Client
	//blurConn    *blur.Blur
}

func New() (*Bot, error) {
	wallet, err := helper.WalletFromPrivateKey(config.Conf.PrivateKey)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(config.Conf.EthNode)
	if err != nil {
		return nil, err
	}

	return &Bot{
		wallet:    wallet,
		ethClient: ethClient,
	}, nil
}

func (b *Bot) Run() {

}
