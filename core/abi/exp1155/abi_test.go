package exp1155

import (
	"fmt"
	"testing"

	"github.com/younamebert/xfssdk/common"
)

func Test_JSON(t *testing.T) {
	objClass, _ := JSON(EXP1155ABI)
	bs, _ := common.MarshalIndent(objClass.Events)
	fmt.Println(string(bs))
}
