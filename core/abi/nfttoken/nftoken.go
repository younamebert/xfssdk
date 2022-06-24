package nfttoken

const (
	CREATE          = "0x0000000000000000000000000000000000000000000000000000000000000000"
	NAME            = "0x1162f326f21ac342307b16730bc30e1cfb6fd35acfd527a2d6adf39d44b56522"
	SYMBOL          = "0xd24b7074b8d5ee3e7e0a471901324f6870e175419253f5e497b42272f6919234"
	BALANCEOF       = "0x61945fbcd9ffbebe7dcf1ec99e8bd195e6b235295dbe5f84df2f8a2b72174e1c"
	OWNEROF         = "0xd00884d5d1e80e12737ac1eeca8780428f2d80e0bb1fda0ed9d9ef9ac460f656"
	ISAPPROVEFORALL = "0x7649e43e1b6b6e86a7dc7f524e8aa0ca2ef18d7ffca9b5d9f92eae1fac6fa36e"
	GETAPPROVED     = "0xed6aa6e08de6aedb04b0bf3cc378a5d7a1e0e339d79f89001feb7377083d7360"

	SETAPPROVAEFORALL = "0x22b4822e0b6ac8a018db0f74c074515486ef1b2b08ac4a285f16c1d4711a014c"
	TRANSFERFROM      = "0x2561555cf5bdc523a9cdcbb7810211f424a3477c8e4ae5773e6a37475247d78a"
	APPROVE           = "0x6007acbe30b2cd98703e83350ea665c06009fcd51f26dd73b309294235f45f21"

	MINT = "0xced97cc4a377b5b4386d9c67bc4f4e14febb561903a27409ce7a2886368b75bb"
)

var NFTOKENABI = `{"events":{"0x0ae9f41b93cbe1d3e94224aa5a1695ed1f9a037995df1dd276ecb309d11135e8":{"name":"NFTokenApprovalEvent","argc":3,"args":[{"name":"Owner","type":"CTypeAddress"},{"name":"Approved","type":"CTypeAddress"},{"name":"TokenId","type":"CTypeUint256"}]},"0x20b529eac07e14d32ddfe96d2f489003c7ca40d0286d2c424e872587c6656805":{"name":"NFTokenTransferEvent","argc":3,"args":[{"name":"From","type":"CTypeAddress"},{"name":"To","type":"CTypeAddress"},{"name":"TokenId","type":"CTypeUint256"}]},"0xe4db419ed5d22abbdb1726213e9ed2cf78c3fc92161b86ff63d56e512449ab11":{"name":"NFTokenApprovalForAllEvent","argc":3,"args":[{"name":"Owner","type":"CTypeAddress"},{"name":"Operator","type":"CTypeAddress"},{"name":"Approved","type":"CTypeBool"}]}},"methods":{"0x0000000000000000000000000000000000000000000000000000000000000000":{"name":"Create","argc":2,"args":[{"name":"","type":"CTypeString"},{"name":"","type":"CTypeString"}],"return_type":"error"},"0x1162f326f21ac342307b16730bc30e1cfb6fd35acfd527a2d6adf39d44b56522":{"name":"GetName","argc":0,"args":[],"return_type":"CTypeString"},"0x22b4822e0b6ac8a018db0f74c074515486ef1b2b08ac4a285f16c1d4711a014c":{"name":"SetApprovalForAll","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeBool"}],"return_type":"CTypeBool"},"0x2561555cf5bdc523a9cdcbb7810211f424a3477c8e4ae5773e6a37475247d78a":{"name":"TransferFrom","argc":3,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0x6007acbe30b2cd98703e83350ea665c06009fcd51f26dd73b309294235f45f21":{"name":"Approve","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0x61945fbcd9ffbebe7dcf1ec99e8bd195e6b235295dbe5f84df2f8a2b72174e1c":{"name":"BalanceOf","argc":1,"args":[{"name":"","type":"CTypeAddress"}],"return_type":"CTypeUint256"},"0x7649e43e1b6b6e86a7dc7f524e8aa0ca2ef18d7ffca9b5d9f92eae1fac6fa36e":{"name":"IsApprovedForAll","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeAddress"}],"return_type":"CTypeBool"},"0xced97cc4a377b5b4386d9c67bc4f4e14febb561903a27409ce7a2886368b75bb":{"name":"Mint","argc":1,"args":[{"name":"","type":"CTypeAddress"}],"return_type":"CTypeUint256"},"0xd00884d5d1e80e12737ac1eeca8780428f2d80e0bb1fda0ed9d9ef9ac460f656":{"name":"OwnerOf","argc":1,"args":[{"name":"","type":"CTypeUint256"}],"return_type":"CTypeAddress"},"0xd24b7074b8d5ee3e7e0a471901324f6870e175419253f5e497b42272f6919234":{"name":"GetSymbol","argc":0,"args":[],"return_type":"CTypeString"},"0xed6aa6e08de6aedb04b0bf3cc378a5d7a1e0e339d79f89001feb7377083d7360":{"name":"GetApproved","argc":1,"args":[{"name":"","type":"CTypeUint256"}],"return_type":"CTypeAddress"}}}`
var NFTBIN = "0xd02302"
