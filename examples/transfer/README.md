## xfssdk-transfer

这个程序演示了一个核心功能交易流程。



### 构建

从xfssdk/examples/transfer目录运行以下命令：

```go
go build || go run
```



### 程序使用和流程结束

**1. 首先创建一笔交易**

```go
handle := xfssdk.Default() // handle对象

	var (
		version    string = "0"//版本
		fromPriKey string = "0x01016ffd70850416510c648c77e7dad721f99dd1d016169f0857716981c963eaf885"//from账户的私钥
		to         string = "nfYXkAZVjZjKnz79RyoLauuAmPBv9DPhi"//交易目标地址
	)

```

