package apis

import "github.com/younamebert/xfssdk/libs/client"

var GVA_XFSCLICENT = new(client.Client)

func SetXFSClient(cli *client.Client) {
	GVA_XFSCLICENT = cli
}
