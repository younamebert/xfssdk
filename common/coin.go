package common

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/younamebert/xfssdk/libs/crypto"
)

const Coin = 10 ^ 6

var AttoCoin = uint64(math.Pow(10, 18))
var NanoCoin = uint64(math.Pow(10, 9))

// gas
func BaseCoin2Atto(coin string) (*big.Int, error) {
	bigCoin, ok := new(big.Float).SetString(coin)
	if !ok {
		return nil, errors.New("string to big.Int Error")
	}
	attocoin := bigCoin.Mul(bigCoin, big.NewFloat(float64(AttoCoin)))
	i, _ := attocoin.Int(nil)
	return i, nil
}

// min coin
func Atto2BaseCoin(atto *big.Int) *big.Int {
	i := big.NewInt(0)
	i.Add(i, atto)
	i.Div(i, big.NewInt(int64(AttoCoin)))
	return i
}

// atto to nano
func AttoCoin2Nano(atto *big.Int) *big.Int {
	i := big.NewInt(0)
	i.Add(i, atto)
	i.Div(i, big.NewInt(int64(NanoCoin)))
	return i
}

// // min coin
// func Atto2BaseCoin(atto *big.Int) *big.Float {
// 	f := new(big.Float).SetInt(atto)
// 	fSub := new(big.Float).SetUint64(AttoCoin)

// 	f.Quo(f, fSub)
// 	return f
// }

func BaseCoin2Nano(coin string) (*big.Int, error) {

	bigCoin, ok := new(big.Float).SetString(coin)
	if !ok {
		return nil, errors.New("string to big.Int Error")
	}
	attocoin := bigCoin.Mul(bigCoin, big.NewFloat(float64(NanoCoin)))
	i, _ := attocoin.Int(nil)
	return i, nil
}

func NanoCoin2BaseCoin(nano *big.Int) *big.Int {
	i := big.NewInt(0)
	i.Add(i, nano)
	i.Div(i, big.NewInt(int64(NanoCoin)))
	return i
}

// nano to atto
func NanoCoin2Atto(nano *big.Int) *big.Int {
	a := new(big.Float).SetInt64(nano.Int64())
	a.Mul(a, big.NewFloat(float64(NanoCoin)))
	i, _ := a.Int(nil)
	return i
}

func Atto2BaseRatCoin(coin string) (*big.Float, error) {
	x, ok := new(big.Float).SetString(coin)
	if !ok {
		return nil, errors.New("string to big.Int Error")
	}
	y := new(big.Float).SetUint64(AttoCoin)
	sum := big.NewFloat(0)
	sum.Quo(x, y)
	return sum, nil
}

func Nano2BaseRatCoin(coin string) (*big.Float, error) {
	x, ok := new(big.Float).SetString(coin)
	if !ok {
		return nil, errors.New("string to big.Int Error")
	}
	y := new(big.Float).SetUint64(NanoCoin)
	sum := big.NewFloat(0)
	sum.Quo(x, y)
	return sum, nil
}

func StrKey2Address(fromprikey string) (Address, error) {

	keyEnc := fromprikey

	if keyEnc[0] == '0' && keyEnc[1] == 'x' {
		keyEnc = keyEnc[2:]
	} else {
		return Address{}, errors.New("binary forward backward error")
	}

	keyDer, err := hex.DecodeString(keyEnc)
	if err != nil {
		return Address{}, err
	}

	_, pKey, err := crypto.DecodePrivateKey(keyDer)
	if err != nil {
		return Address{}, err
	}

	addr := crypto.DefaultPubKey2Addr(pKey.PublicKey)
	return addr, nil
}

func CheckWalletPriKey(key string) error {

	keyEnc := key

	if keyEnc[0] == '0' && keyEnc[1] == 'x' {
		keyEnc = keyEnc[2:]
	} else {
		return errors.New("binary forward backward error")
	}

	keyDer, err := hex.DecodeString(keyEnc)
	if err != nil {
		return err
	}
	kv, _, err := crypto.DecodePrivateKey(keyDer)
	if err != nil {
		return err
	}
	if kv != crypto.DefaultKeyPackVersion {
		return fmt.Errorf("unknown private key version %d", kv)
	}
	return nil
}
