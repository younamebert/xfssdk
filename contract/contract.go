package contract

import (
	"github.com/younamebert/xfssdk/contract/nfttoken"
	"github.com/younamebert/xfssdk/contract/stdtoken"
)

type ContractEngine struct {
	StdToken stdtoken.StdTokenCall
	NFTToken nfttoken.NFTTokenCall
}
