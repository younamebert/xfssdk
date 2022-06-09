package apicontract

import (
	"testing"

	"github.com/younamebert/xfssdk/core/apis"
)

var testTokenObj = new(ApiToken)

func Test_Create(t *testing.T) {
	testdata := TokenArgs{
		Name:        "sdk",
		Symbol:      "sdkxfs",
		Decimals:    "10000000000",
		TotalSupply: "100000000000000",
	}
	if err := apis.XFSABI(); err != nil {
		t.Fatal(err)
		return
	}
	code, err := testTokenObj.Create(testdata)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("code: %v", code)

	// testdeploy := DeployTokenArgs{

	// }
}
