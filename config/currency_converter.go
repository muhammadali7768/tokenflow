package config

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

func ConvertEtherToWei(etherAmount float64) (*big.Int, error) {
	// Parse the float to a big.Float
	etherValue := new(big.Float).SetFloat64(etherAmount)

	// Convert Ether to Wei (1 Ether = 1e18 Wei)
	weiValue := new(big.Float).Mul(etherValue, big.NewFloat(params.Ether))

	// Convert to big.Int for transaction use
	wei := new(big.Int)
	weiValue.Int(wei)

	return wei, nil
}
func ConvertWeiToEther(weiAmount *big.Int) *big.Float {
	weiAsFloat := new(big.Float).SetInt(weiAmount)
	etherConversionFactor := new(big.Float).SetInt(big.NewInt(1000000000000000000))
	etherAmount := new(big.Float).Quo(weiAsFloat, etherConversionFactor)
	return etherAmount
}
