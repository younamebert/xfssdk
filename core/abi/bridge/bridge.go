package bridge

const (
	CREATE       = "0x0000000000000000000000000000000000000000000000000000000000000000"
	TANSFERIN    = "0x40d2e803bbe3904840cece9e5997f40fa8cc16db20357606ddb096193dcdcdce"
	TANSFERINOUT = "0xb6b0adfd06b688ce0bebc64dfdeefdddbc38a71a7ca0de0869499668245e41d4"
)

var BRIDGETOKENABI = `{"events":{"0x06976fbcd8c1677b395e07d3c109af81384fea5da3dbce252de69d194c97d4bb":{"name":"BridgeTransferInEvent","argc":4,"args":[{"name":"BankAddress","type":"CTypeAddress"},{"name":"DepositorAddress","type":"CTypeAddress"},{"name":"ContractAddress","type":"CTypeAddress"},{"name":"Value","type":"CTypeUint256"}]},"0x0db84af94380c5f8b90cb557eabb4534c01d07dd3775ba4ab0d91cb94ec898a6":{"name":"BridgeTransferOutEvent","argc":4,"args":[{"name":"BankAddress","type":"CTypeAddress"},{"name":"DepositorAddress","type":"CTypeAddress"},{"name":"ContractAddress","type":"CTypeAddress"},{"name":"Value","type":"CTypeUint256"}]}},"methods":{"0x0000000000000000000000000000000000000000000000000000000000000000":{"name":"Create","argc":4,"args":[{"name":"","type":"CTypeString"},{"name":"","type":"CTypeString"},{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0x40d2e803bbe3904840cece9e5997f40fa8cc16db20357606ddb096193dcdcdce":{"name":"TransferIn","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0xb6b0adfd06b688ce0bebc64dfdeefdddbc38a71a7ca0de0869499668245e41d4":{"name":"TransferOut","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0xd4e97969f51b509b07924c476ae8983687162f7e7e5068603ca2780ebf18451b":{"name":"GetChainId","argc":0,"args":[],"return_type":"CTypeUint256"}}}`
var BRIDGEBIN = "0xd02303"
