package crypto

import (
	"encoding/hex"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/common/ahash"
)

func GetCAddr(address common.Address, nonce uint64) common.Address {

	fromAddressHashBytes := ahash.SHA256(address[:])
	fromAddressHash := common.Bytes2Hash(fromAddressHashBytes)
	caddr := CreateAddress(fromAddressHash, nonce)
	return caddr
}

func Prikey2Addr(key string) common.Address {
	der, err := hex.DecodeString(key)
	if err != nil {
		return common.ZeroAddr
	}
	_, pKey, err := DecodePrivateKey(der)
	if err != nil {
		return common.ZeroAddr
	}
	return DefaultPubKey2Addr(pKey.PublicKey)
}
