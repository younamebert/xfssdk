package nfttoken

import (
	"fmt"
	"testing"

	"github.com/younamebert/xfssdk/common"
)

func Test_JSON(t *testing.T) {
	objClass, _ := JSON(NFTOKENABI)
	bs, _ := common.MarshalIndent(objClass.Events)
	fmt.Println(string(bs))
}
