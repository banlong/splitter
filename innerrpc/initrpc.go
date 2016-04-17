package innerrpc
//THIS MODULE INITIALIZES RPC CONNECTION TO DATA PROVIDE

import (
	"time"
	"splitter/config"
)

var (
	DataClient   *Client
	err error
	dns = config.DATASERVER_IP + config.DBSERVER_RPC_PORT

)

func Start() {
	//Dial Data Provider
	DataClient, err = NewClient(dns, time.Millisecond*500)
}



