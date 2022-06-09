package common

import "math/big"

var TxGas = big.NewInt(25000)
var TxGasPrice = big.NewInt(10)

func DefaultGasPrice() *big.Int {
	return NanoCoin2Atto(TxGasPrice)
}
