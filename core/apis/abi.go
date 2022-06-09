package apis

import (
	"github.com/younamebert/xfssdk/core/abi"
)

var GVA_ABI abi.ABI

// SetXFSClient set API global request client
func XFSABI() error {
	abi, err := abi.JSON(abi.XFSTOKENABI)
	if err != nil {
		return err
	}
	GVA_ABI = abi
	return nil
}
