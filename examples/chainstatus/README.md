## xfssdk-chainstatus

这个程序演示了一个获取节点信息流程。

### 构建

从xfssdk/examples/chainstatus目录运行以下命令：

```go
go build || go run
```

### 程序使用和流程

```go
// 初始化客户端
cli := client.NewClient("https://api.scan.xfs.tech/jsonrpc/v2/", "5s")
// 设置全局客户端请求
apis.SetXFSClient(cli)

apimethod := api.NewApiMethod()
//查询节点最新的blockheader信息
latestBlockHeader, err := apimethod.Chain.GetHead()
```

### 结果

```json
{
    "height": 34339,
    "version": 0,
    "hash_prev_block": "0x000000016466d7c6dab40b760f9f3f6e96762a54cd9c95f68beab018000e9a2a",
    "timestamp": 1647326354,
    "coinbase": "edarvPH49waieoqT2mS9KqL3hvGTLXzxt",
    "state_root": "0x49b8b6478aeea95637ad0154e7ff387c71208378364d33a0d29f00f25f92da06",
    "transactions_root": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "receipts_root": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "gas_limit": 25000000,
    "gas_used": 0,
    "bits": 178534941,
    "nonce": 90231242,
    "extranonce": 11019934332299970000,
    "hash": "0x00000009f00b3cb03141426ab98b010f3b6f291771d2e93952be15f8026f97f8"
}
```

