package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/younamebert/xfssdk"
	"github.com/younamebert/xfssdk/exactly/inspecttx"
)

func main() {
	handle := xfssdk.Default()

	var (
		version    string = "0"
		fromPriKey string = "0x01016ffd70850416510c648c77e7dad721f99dd1d016169f0857716981c963eaf885"
		to         string = "nfYXkAZVjZjKnz79RyoLauuAmPBv9DPhi"
	)
	//私钥映射出address
	address, err := handle.Exactly.InspectTx.GetFromAddress(fromPriKey)
	if err != nil {
		fmt.Printf("prikey to address err:%v\n", err)
		os.Exit(1)
	}
	// 获取from最新在交易池的nonce值
	nonce, err := handle.ApiMethod.TxPool.GetAddrTxNonce(address.B58String())
	if err != nil {
		fmt.Printf("by address get txpool nonce err:%v\n", err)
		os.Exit(1)
	}
	noncestr := strconv.FormatInt(*nonce, 10)
	tx := inspecttx.StringRawTransaction{
		Version:  version,
		To:       to,
		GasPrice: "10000000000",
		GasLimit: "25000",
		Nonce:    noncestr,
		Value:    "1",
	}

	//交易签名
	if err := tx.SignWithPrivateKey(fromPriKey); err != nil {
		fmt.Printf("by address get txpool nonce err:%v\n", err)
		os.Exit(1)
	}

	// 把交易对象结构体base64加密
	txraw, err := tx.RawTx()
	if err != nil {
		fmt.Printf("tx to base64 err:%v\n", err)
		os.Exit(1)
	}
	// 发送一笔交易
	txhash, err := handle.ApiMethod.TxPool.SendRawTransaction(txraw)
	if err != nil {
		fmt.Printf("ApiMethod txpool  SendRawTransaction err:%v\n", err)
		os.Exit(1)
	}
	fmt.Println(txhash)
}
