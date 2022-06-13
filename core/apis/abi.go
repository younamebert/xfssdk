package apis

import (
	nfttokenabi "github.com/younamebert/xfssdk/core/abi/nfttoken"
	stdtokenabi "github.com/younamebert/xfssdk/core/abi/stdtoken"
)

var (
	GVA_ABI_STDTOKEN stdtokenabi.ABI
	GVA_ABI_NFTTOKEN nfttokenabi.ABI
)

// SetXFSClient set API global request client
func XFSABI() error {
	stdabi, err := stdtokenabi.JSON(stdtokenabi.XFSTOKENABI)
	if err != nil {
		return err
	}
	GVA_ABI_STDTOKEN = stdabi

	nftabi, err := nfttokenabi.JSON(nfttokenabi.NFTOKENABI)
	if err != nil {
		return err
	}
	GVA_ABI_NFTTOKEN = nftabi
	return nil
}
