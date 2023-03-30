package utils

import (
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
)

const ethToWei = float64(1000000000000000000)

func EtherToWei(n float64) *big.Int {
	return floatToBigInt(n, 18)
}

func WeiToEther(wei *big.Int) float64 {
	eth := big.NewInt(1000000000000000000)
	if wei.Cmp(eth) == 1 {
		var result big.Int
		result.Div(wei, eth)
		return float64(result.Int64())
	} else {
		return float64(wei.Uint64()) / ethToWei
	}
}

func WeiStringToEther(weiStr string) float64 {
	// convert Wei string to big.Int
	wei := new(big.Int)
	if _, ok := wei.SetString(weiStr, 10); !ok {
		return 0
	}

	// create big.Float with wei value
	ethFloat := new(big.Float).SetInt(wei)

	// calculate ETH value by dividing by 10^18
	ethFloat.Quo(ethFloat, big.NewFloat(1e18))

	// convert big.Float to float64
	eth, _ := ethFloat.Float64()

	return eth
}

func floatToBigInt(amount float64, decimal int64) *big.Int {
	// 6 is our smallest precision
	if decimal < 6 {
		return big.NewInt(int64(amount * math.Pow10(int(decimal))))
	}
	result := big.NewInt(int64(amount * math.Pow10(6)))
	return result.Mul(result, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(decimal-6), nil))
}

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
