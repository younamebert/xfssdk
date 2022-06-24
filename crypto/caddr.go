package crypto

import (
	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/common/ahash"
)

func GetCAddr(address common.Address, nonce uint64) common.Address {

	fromAddressHashBytes := ahash.SHA256(address[:])
	fromAddressHash := common.Bytes2Hash(fromAddressHashBytes)
	caddr := CreateAddress(fromAddressHash, nonce)
	return caddr
}
