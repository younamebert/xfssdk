package checkprikey

import (
	"fmt"
	"xfssdk/libs/crypto"
)

type CheckPriKeyWay interface {
	CheckWalletPriKey(der []byte) error
}

type CheckPriKey struct {
}

func (checkprikey *CheckPriKey) CheckWalletPriKey(der []byte) error {
	kv, _, err := crypto.DecodePrivateKey(der)
	if err != nil {
		return err
	}
	if kv != crypto.DefaultKeyPackVersion {
		return fmt.Errorf("unknown private key version %d", kv)
	}
	return nil
}
