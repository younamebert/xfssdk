package apis

import (
	bridgeabi "github.com/younamebert/xfssdk/core/abi/bridge"
	nftmarketabi "github.com/younamebert/xfssdk/core/abi/nftmarket"
	nfttokenabi "github.com/younamebert/xfssdk/core/abi/nfttoken"
	stdtokenabi "github.com/younamebert/xfssdk/core/abi/stdtoken"
)

var (
	GVA_ABI_STDTOKEN    stdtokenabi.ABI
	GVA_ABI_NFTTOKEN    nfttokenabi.ABI
	GVA_ABI_BRIDGETOKEN bridgeabi.ABI
	GVA_ABI_NFTMARKET   nftmarketabi.ABI
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

	bridgeabi, err := bridgeabi.JSON(bridgeabi.BRIDGETOKENABI)
	if err != nil {
		return err
	}
	GVA_ABI_BRIDGETOKEN = bridgeabi

	nftmarketabi, err := nftmarketabi.JSON(nftmarketabi.NFTMARKETABI)
	if err != nil {
		return err
	}
	GVA_ABI_NFTMARKET = nftmarketabi
	return nil
}
