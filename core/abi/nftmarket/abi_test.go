package nftmarket

import (
	"fmt"
	"testing"

	"github.com/younamebert/xfssdk/common"
)

func Test_JSON(t *testing.T) {
	objClass, _ := JSON(NFTMARKETABI)
	bs, _ := common.MarshalIndent(objClass.Events)
	fmt.Println(string(bs))
}
