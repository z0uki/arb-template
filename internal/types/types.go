package types

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

type TxnData struct {
	To          string
	Data        string
	Value       *big.Int
	GasEstimate int64
}

type MatchedOrders struct {
	//matched orders info
}
