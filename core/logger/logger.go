package logger

import "xfssdk/libs/client"

var XFSCLICENT = new(client.Client)

func SetXFSClient(cli *client.Client) {
	XFSCLICENT = cli
}
