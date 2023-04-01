package helper

import (
	"log"
	"math"
	"math/big"
	"time"

	"context"

	"arb-template/internal/types"
	"arb-template/internal/utils"

	"github.com/0xblocks/flashbots-bundle"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Flashbots struct {
	Provider *flashbots.Provider
	EthNode  *ethclient.Client
	Wallet   *types.Wallet

	chainID *big.Int
}

func NewFlashbots(ethNode *ethclient.Client, wallet *types.Wallet, signWallet *types.Wallet) (*Flashbots, error) {

	chainID, err := ethNode.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	relayUrl := flashbots.DefaultRelayURL
	if chainID.Cmp(big.NewInt(1)) != 0 {
		relayUrl = flashbots.TestRelayURL
	}

	provider := flashbots.NewProvider(signWallet.PrivateKey, wallet.PrivateKey, relayUrl)

	return &Flashbots{
		Provider: provider,
		EthNode:  ethNode,
		Wallet:   wallet,
		chainID:  chainID,
	}, nil
}

func (f *Flashbots) BuildTransaction(txn *types.TxnData, nonce uint64) (*geth_types.Transaction, error) {
	var (
		err       error
		value     = txn.Value
		data      = common.FromHex(txn.Data)
		to        = common.HexToAddress(txn.To)
		gasLimit  = uint64(500000)
		gasFeeCap = big.NewInt(10000000000) //10gwei
		gasTipCap = big.NewInt(2000000000)  //2gwei
	)

	if nonce == 0 {
		nonce, err = f.EthNode.PendingNonceAt(context.Background(), f.Wallet.Address)
		if err != nil {
			return nil, err
		}
	}

	if txn.GasEstimate != 0 {
		gasLimit = uint64(txn.GasEstimate)
	}

	gasFeeCap, err = f.EthNode.SuggestGasPrice(context.Background())
	if gasFeeCap.Cmp(gasTipCap) < 0 {
		gasFeeCap = gasFeeCap.Add(gasFeeCap, gasTipCap)
	}

	tx := geth_types.NewTx(&geth_types.DynamicFeeTx{
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		To:        &to,
		Value:     value,
		Data:      data,
	})

	return geth_types.SignTx(tx, geth_types.LatestSignerForChainID(f.chainID), f.Wallet.PrivateKey)
}

func (f *Flashbots) SendBundle(signedTxs []*geth_types.Transaction) error {
	start := time.Now()
	attempts := 1
	if f.chainID.Cmp(big.NewInt(1)) != 0 {
		//Testnet
		attempts = 10
	}

	txs := []string{}
	for _, tx := range signedTxs {
		data, err := tx.MarshalBinary()
		if err != nil {
			return err
		}

		txs = append(txs, hexutil.Encode(data))
	}

	block, err := f.EthNode.BlockByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	//First simulate the transaction to make sure it doesn't revert
	resp, err := f.Provider.Simulate(txs, block.Number(), "latest", 0)
	if err != nil {
		return err
	}

	if err = resp.HasError(); err != nil {
		return err
	}

	coinbase, _ := new(big.Float).SetString(resp.Result.CoinbaseDiff)
	eth := new(big.Float).Quo(coinbase, big.NewFloat(math.Pow10(18)))
	wei, _ := resp.EffectiveGasPrice()
	gwei := utils.EthToGwei(wei)

	log.Printf("Simulation completed in %fs. Cost: %f eth, effective price: %d gwei\n", time.Since(start).Seconds(), eth, gwei)

	// Send the bundle to the FB relay for inclusion in a block
	// Unless your tx is only valid for a single block, you should target several blocks since flashbots
	// isn't available on 100% of the hashing power. On Goerli, it's only available on a single validator,
	// so target at least 9 blocks to ensure it gets picked up
	for i := 0; i < attempts; i++ {
		targetBlockNumber := new(big.Int).Add(block.Number(), big.NewInt(int64(i)))
		if _, err = f.Provider.SendBundle(txs, targetBlockNumber, &flashbots.Options{}); err != nil {
			return err
		}

		log.Printf("submitted for block: %d\n", targetBlockNumber)
	}
	return nil
}

func (f *Flashbots) WaitTx(tx *geth_types.Transaction, maxWaitSeconds uint) error {
	log.Println("Waiting for tx to complete...")
	loops := uint(0)

	for {
		_, isPending, err := f.EthNode.TransactionByHash(context.Background(), tx.Hash())
		if err != nil && err != ethereum.NotFound {
			return err
		}

		if !isPending {
			// It's not pending, so check if it's been mined
			receipt, err := f.EthNode.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil && err != ethereum.NotFound {
				return err
			}
			if receipt != nil {
				log.Println("tx complete. it may take a few minutes to appear in etherscan.")
				if f.chainID.Cmp(big.NewInt(1)) != 0 {
					log.Printf("https://goerli.etherscan.io/tx/%s\n", tx.Hash().Hex())
				} else {
					log.Printf("https://etherscan.io/tx/%s\n", tx.Hash().Hex())
				}
				break
			}
		}

		time.Sleep(1 * time.Second)

		loops = loops + 1
		if loops > maxWaitSeconds {
			log.Printf("timed out after %d seconds. check manually here:\n", maxWaitSeconds)
			if f.chainID.Cmp(big.NewInt(1)) != 0 {
				log.Printf("https://goerli.etherscan.io/tx/%s\n", tx.Hash().Hex())
			} else {
				log.Printf("https://etherscan.io/tx/%s\n", tx.Hash().Hex())
			}
			break
		}
	}

	return nil
}

// CheckAndBribe Add as needed
func (f *Flashbots) CheckAndBribe(bribe *big.Int) error {
	return nil
}
