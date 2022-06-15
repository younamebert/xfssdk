package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"

// 	"github.com/younamebert/xfssdk"
// 	"github.com/younamebert/xfssdk/exactly/inspecttx"
// )

// func main() {
// 	handle := xfssdk.Default()

// 	var (
// 		version    string = "0"
// 		fromPriKey string = "0x01016ffd70850416510c648c77e7dad721f99dd1d016169f0857716981c963eaf885"
// 		to         string = "nfYXkAZVjZjKnz79RyoLauuAmPBv9DPhi"
// 	)
// 	// the private key is mapped to the account address
// 	address, err := handle.Exactly.InspectTx.GetFromAddress(fromPriKey)
// 	if err != nil {
// 		fmt.Printf("prikey to address err:%v\n", err)
// 		os.Exit(1)
// 	}
// 	// Get the nonce value of the latest from in the trading pool
// 	nonce, err := handle.ApiMethod.TxPool.GetAddrTxNonce(address.B58String())
// 	if err != nil {
// 		fmt.Printf("by address get txpool nonce err:%v\n", err)
// 		os.Exit(1)
// 	}
// 	noncestr := strconv.FormatInt(*nonce, 10)
// 	tx := inspecttx.StringRawTransaction{
// 		Version:  version,
// 		To:       to,
// 		GasPrice: "10000000000",
// 		GasLimit: "25000",
// 		Nonce:    noncestr,
// 		Value:    "1",
// 	}

// 	// transaction signature
// 	if err := tx.SignWithPrivateKey(fromPriKey); err != nil {
// 		fmt.Printf("by address get txpool nonce err:%v\n", err)
// 		os.Exit(1)
// 	}

// 	// Encrypt the transaction object structure Base64
// 	txraw, err := tx.Transfer2Raw()
// 	fmt.Println(txraw)
// 	if err != nil {
// 		fmt.Printf("tx to base64 err:%v\n", err)
// 		os.Exit(1)
// 	}
// 	// send a transaction
// 	txhash, err := handle.ApiMethod.TxPool.SendRawTransaction(txraw)
// 	if err != nil {
// 		fmt.Printf("ApiMethod txpool  SendRawTransaction err:%v\n", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(txhash)
// }
