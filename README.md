## xfssdk-go

**xfssdk** it is a class library that can link to xfscoin. hrough some configurations, you can obtain the chain status, send transactions, etc. the library is intended to be friendly to developers	

### Contents

- **apichain**
  - **apitxpool**
  - **apistate**
  - **apichain**
  - **apinet**
- **exactly**
  - **inspecttx**
  - **checkprikey**

### Directory structure description

```
├── README.md   // help
├── api         // chain method status
├── common      // public Toolkit 
├── config      // handleconfig & loggerconfig
├── core            
│   └── apis    // global call return of API
│   └── logger.go
├── exactly 
│   └── checkprikey // private key verification
│   └── inspecttx  // generate transaction signature and send
├── examples
│   └── chainstatus // examples chainstatus
│   └── transfer    // examples transfer
├── global
├── libs
│   └── ahash  // encode
│   └── client // xfscoin client service dependency
│   └── crypto // 256k1
├── utils
├── handle.go // entry service file
├── go.mod
├── go.sum
```



### Installation

To install **xfssdk** package, you need to install Go and set your Go workspace first.

**1. First you need to install[Go （need](https://golang.org/)**1.13+ version**),then you can use the below Go command to install xfssdk**

```shell
$ go get -u github.com/younamebert/xfssdk
```

**2. Import it in your code:**

```go
import "github.com/younamebert/xfssdk"
```

### Quick start

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

### Other

1. godoc

   ```shell
   $ godoc  -goroot=. -http=:6061
   ```

### Example

in [sample folder](https://github.com/younamebert/xfssdk/tree/main/examples)you can find a way to send a transaction or get the chain status to help you get started**xfssdk**。

### Contribution

If you want to contribute to **xfssdk** please fork, repair, submit and send a pull request. Submissions that do not meet the coding criteria will be ignored (using gofmt!). If you send a pull request, make sure that you submit on the development branch and that you do not merge into the main branch. Submissions based directly on main are ignored.

### License

Issued under [MIT license](https://github.com/go-gorm/gorm/blob/master/License) And provided "as is" without any express or implied warranty. Any security provided by XFS software depends in part on its use, configuration and deployment **xfssdk** is based on many third-party libraries, XFS Tech does not represent or warrant that **xfssdk** or any third-party library will perform as expected or will not have errors, errors or error codes. Both may fail in large or small ways, thus completely or partially limiting the function or endangering the computer system. If you use or implement **xfssdk**, you will do so at your own risk. In any case, XFS Tech shall not be liable to any party for any damages, even if it has been informed of the possibility of damage.

