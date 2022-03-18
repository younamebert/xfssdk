## xfssdk-chainstatus

this program demonstrates a process of obtaining node information.

### Build

From the xfssdk/examples/chainstatus directory run the following：

```go
go build || go run
```

### application and process

```go
// Initialize client
cli := client.NewClient("https://api.scan.xfs.tech/jsonrpc/v2/", "5s")
// Set global client requests
apis.SetXFSClient(cli)

apimethod := api.NewApiMethod()
// Query the latest blockheader information of the node
latestBlockHeader, err := apimethod.Chain.GetHead()
```

### result

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

