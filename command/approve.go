package command

import (
	"arb-template/internal/types"
	"arb-template/internal/utils"
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"

	"arb-template/internal/binds/bind_weth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	NodeURL           = "https://geth.mytokenpocket.vip"
	WETHAddress       = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	OpenSeaFeeAddress = common.HexToAddress("0x1E0049783F008A0085193E00003D00cd54003c71")
)

func newWethInstance() *bind_weth.Weth {
	contractAddress := common.HexToAddress(WETHAddress)
	client, err := ethclient.Dial("https://geth.mytokenpocket.vip")
	if err != nil {
		log.Println(err)
		return nil
	}
	instance, err := bind_weth.NewWeth(contractAddress, client)
	if err != nil {
		log.Println(err)
		return nil
	}
	return instance
}

func approveWETH(wallet *types.Wallet) error {
	w := newWethInstance()

	allowance, _ := w.Allowance(nil, wallet.Address, OpenSeaFeeAddress)

	if allowance.String() != "0" {
		return errors.New("address: " + wallet.Address.Hex() + " has approved!")
	}

	Opts, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, big.NewInt(1))
	if err != nil {
		return err
	}

	tx, err := w.Approve(Opts, OpenSeaFeeAddress, utils.EthToWei(1000))
	if err != nil {
		return err
	}

	return waitMined(tx)
}

func waitMined(tx *geth_types.Transaction) error {
	ec, err := ethclient.Dial(NodeURL)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(context.Background(), ec, tx)
	return err
}
