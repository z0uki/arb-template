package helper

import (
	"crypto/ecdsa"
	"fmt"

	"arb-template/internal/types"

	"github.com/ethereum/go-ethereum/crypto"
)

func WalletFromPrivateKey(pkey string) (*types.Wallet, error) {
	privateKey, err := crypto.HexToECDSA(pkey)
	if err != nil {
		return nil, err
	}

	// get public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// get address
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &types.Wallet{
		PrivateKey: privateKey,
		Address:    address,
	}, nil
}
