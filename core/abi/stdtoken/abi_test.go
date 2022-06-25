package stdtoken

import (
	"fmt"
	"testing"

	"github.com/younamebert/xfssdk/common"
)

func Test_JSON(t *testing.T) {
	objClass, _ := JSON(XFSTOKENABI)
	bs, _ := common.MarshalIndent(objClass.Events)
	fmt.Println(string(bs))
}

func Test_PackEventsName(t *testing.T) {
	abi, err := JSON(XFSTOKENABI)
	if err != nil {
		t.Fatal(err)
		return
	}
	testStr := `{"from":"00000000000000000000000000000000000000000000000000","to":"01796e28058b703693d4c786f2b5d408706316364132acc382","value":"00000000000000000000000000000000000000000000003635c9adc5dea00000"}`

	events, err := Str2Events(testStr)
	if err != nil {
		t.Fatal(err)
		return
	}
	eventsResp, err := abi.PackEventsName(events)
	if err != nil {
		t.Fatal(err)
		return
	}
	bs, err := common.MarshalIndent(eventsResp)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(string(bs))
}
