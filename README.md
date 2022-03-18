## xfssdk-go

**xfssdk**是一个可以链接请求**xfscoin**的类库，通过某些配置，可以获取链状态，发送交易等等，库旨对开发人员友好。

### 模块概述

- **apichain链状态模块**
  - **txpool交易池**
  - **state账户状态**
  - **chain状态**
  - **net网络状态**
- **exactly模块**
  - **inspecttx交易操作**
  - **checkprikey签名验证**

### 安装

要安装 **xfssdk**包，您需要先安装 Go 并设置您的 Go 工作区。

**1. 首先需要安装[Go （需要](https://golang.org/)**1.13+ 版本**），然后你可以使用下面的 Go 命令安装 xfssdk**

```shell
$ go get -u github.com/younamebert/xfssdk
```

**2. 在您的代码中导入： **

```go
import "github.com/younamebert/xfssdk"
```

### 快速入门

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/younamebert/xfssdk"
)

func main() {
	handle := xfssdk.Default()
	latestBlockHeader, err := handle.ApiMethod.Chain.GetHead()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	bs, err := json.MarshalIndent(latestBlockHeader, "  ", "")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(bs))
}

```

### 其他

1. 支持godoc

   ```shell
   $ godoc  -goroot=. -http=:6061
   ```

## 例子

在[示例文件夹](https://github.com/younamebert/xfssdk/tree/main/examples)中，您可以找到一个发送一笔交易或者获取链状态来帮助您开始使用**xfssdk**。

### 贡献

如果您想为 **xfssdk** 做出贡献，请分叉、修复、提交并发送拉取请求。不符合编码标准的提交将被忽略（使用 gofmt!)。如果您发送拉取请求，请绝对确保您在开发分支上提交并且您不会合并到主分支。直接基于 main 的提交会被忽略。

### 执照

在[MIT 许可下发布](https://github.com/go-gorm/gorm/blob/master/License),并且"按原样"提供，没有任何明示或暗示的保证。XFS 软件提供的任何安全性部分取决于它的使用、配置和部署方式。**xfssdk** 建立在许多第三方库的基础上，xfs.tech 不声明或保证 **xfssdk** 或任何第三方库将按预期执行或不会出现错误、错误或错误代码。两者都可能以或大或小的方式失败，从而完全或部分限制功能或危及计算机系统。如果您使用或实施 **xfssdk**，您将自行承担风险。在任何情况下，xfs.tech 均不对任何一方承担任何损害赔偿责任，即使已被告知存在损害的可能性。

