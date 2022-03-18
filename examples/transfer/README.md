## xfssdk-transfer

这个程序演示了一个发送交易流程。

### 构建

从xfssdk/examples/transfer目录运行以下命令：

```go
go build || go run
```

### 程序使用和流程

**1. 首先创建一笔交易**

```go

	// 私钥映射出address
	handle.Exactly.InspectTx.GetFromAddress(fromPriKey)
	// 获取from最新在交易池的nonce值
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

**2. 生成交易签名 **

```go
//交易签名
tx.SignWithPrivateKey(fromPriKey);
//把交易对象结构体base64加密
tx.RawTx()
```

**3. 发送一笔交易**

```go
handle.ApiMethod.TxPool.SendRawTransaction(txraw)
```

### 结果

```
0x34be6a97948c91553c964b0ff7160622a2d0d1db9083cdb0558f9a546d7ece93
```

