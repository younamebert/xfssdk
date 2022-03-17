package exactly

import (
	"github.com/younamebert/xfssdk/exactly/checkprikey"
	"github.com/younamebert/xfssdk/exactly/inspecttx"
)

type Exactly struct {
	InspectTx   inspecttx.InspectTxWay
	CheckPriKey checkprikey.CheckPriKeyWay
}

func NewExactly() *Exactly {
	return &Exactly{
		InspectTx:   new(inspecttx.InspectTx),
		CheckPriKey: new(checkprikey.CheckPriKey),
	}
}
