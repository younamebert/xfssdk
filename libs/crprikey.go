package libs

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/crypto"
)

func StrKey2Address(fromprikey string) (common.Address, error) {

	keyEnc := fromprikey

	if keyEnc[0] == '0' && keyEnc[1] == 'x' {
		keyEnc = keyEnc[2:]
	} else {
		return common.Address{}, errors.New("binary forward backward error")
	}

	keyDer, err := hex.DecodeString(keyEnc)
	if err != nil {
		return common.Address{}, err
	}

	_, pKey, err := crypto.DecodePrivateKey(keyDer)
	if err != nil {
		return common.Address{}, err
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
