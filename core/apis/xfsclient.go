package apis

import "github.com/younamebert/xfssdk/libs/client"

var GVA_XFSCLICENT = new(client.Client)

// SetXFSClient 设置api全局请求client
func SetXFSClient(cli *client.Client) {
	GVA_XFSCLICENT = cli
}
