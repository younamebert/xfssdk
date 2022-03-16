package exactly

import (
	"xfssdk/exactly/checkprikey"
	"xfssdk/exactly/inspecttx"
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
