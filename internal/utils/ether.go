package utils

import (
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
)

const ethToWei = float64(1000000000000000000)

// EthToWei convert eth to wei
func EthToWei(n float64) *big.Int {
	return floatToBigInt(n, 18)
}

// EthToGwei convert eth to gwei
func EthToGwei(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, big.NewInt(1000000000))
}

// WeiToEth convert wei to eth
func WeiToEth(wei *big.Int) float64 {
	eth := big.NewInt(1000000000000000000)
	if wei.Cmp(eth) == 1 {
		var result big.Int
		result.Div(wei, eth)
		return float64(result.Int64())
	} else {
		return float64(wei.Uint64()) / ethToWei
	}
}

// WeiToEthS convert wei to eth string
func WeiToEthS(weiStr string) float64 {
	wei := new(big.Int)
	if _, ok := wei.SetString(weiStr, 10); !ok {
		return 0
	}

	ethFloat := new(big.Float).SetInt(wei)

	ethFloat.Quo(ethFloat, big.NewFloat(1e18))

	eth, _ := ethFloat.Float64()

	return eth
}

// StringToBigInt convert string to big.Int
func StringToBigInt(s string) *big.Int {
	var base = 10
	if strings.Contains(s, "0x") {
		s = strings.Replace(s, "0x", "", 1)
		base = 16
	} else if strings.ContainsAny(s, "eE") {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Println(s, "cannot convert string to big.Int")
			return big.NewInt(0)
		}
		i := new(big.Int)
		i.SetString(strconv.FormatFloat(f, 'f', -1, 64), 10)
		return i
	}

	i, ok := new(big.Int).SetString(s, base)
	if !ok {
		log.Println(s, "cannot convert string to big.Int")
		return big.NewInt(0)
	}
	return i
}

func floatToBigInt(amount float64, decimal int64) *big.Int {
	if decimal < 6 {
		return big.NewInt(int64(amount * math.Pow10(int(decimal))))
	}
	result := big.NewInt(int64(amount * math.Pow10(6)))
	return result.Mul(result, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(decimal-6), nil))
}
