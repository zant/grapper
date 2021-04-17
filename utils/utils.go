package utils

import (
  "math"
  "math/big"
)

func WeiToEth(wei *big.Int) *big.Float {
  fbalance := new(big.Float)
  fbalance.SetString(wei.String())
  ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
  return ethValue
}
