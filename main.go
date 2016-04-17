package main
import (
	"splitter/edgerpc"
	"splitter/innerrpc"
)

func main()  {

	//Start RPC connect to DATA PROVIDER
	go innerrpc.Start()

	//Start RPC connect to UPLOAD MASTER
	edgerpc.Start()




}


