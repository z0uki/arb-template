package types

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}
