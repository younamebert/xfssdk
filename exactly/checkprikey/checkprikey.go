package checkprikey

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/younamebert/xfssdk/libs/crypto"
)

type CheckPriKeyWay interface {
	CheckWalletPriKey(der string) error
}

type CheckPriKey struct {
}

func (checkprikey *CheckPriKey) CheckWalletPriKey(der string) error {

	keyEnc := der

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
