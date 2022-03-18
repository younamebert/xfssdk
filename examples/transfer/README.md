## xfssdk-transfer

this program demonstrates a sending transaction process.

### Build

From the xfssdk/examples/transfer run the following：

```go
go build || go run
```

### application and process

**1. First create a transaction**

```go

	// The private key is mapped to the account address
	handle.Exactly.InspectTx.GetFromAddress(fromPriKey)
	// Get the nonce value of the latest from in the trading pool
	handle.ApiMethod.TxPool.GetAddrTxNonce(address.B58String())
tx := inspecttx.StringRawTransaction{
		Version:  version,
		To:       to,
		GasPrice: "10000000000",
		GasLimit: "25000",
		Nonce:    noncestr,
		Value:    "1",
}
```

**2. generate transaction signature**

```go
// Transaction signature
tx.SignWithPrivateKey(fromPriKey);
// Encrypt the transaction object structure Base64
tx.RawTx()
```

**3. send transaction**

```go
handle.ApiMethod.TxPool.SendRawTransaction(txraw)
```

### result

```
0x34be6a97948c91553c964b0ff7160622a2d0d1db9083cdb0558f9a546d7ece93
```

