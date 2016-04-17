package edgerpc

import (
	"net"
	"log"
	"time"
	"splitter/config"
	"splitter/tools"
	"splitter/lib/cenkalti/rpc2"
)

// This will start a RPC connection with UploadMaster
// PROVIDES THE UPLOAD API THAT WILL BE CALL FROM THE UPLOAD MASTER
func Start()  {

	//temp dir for processing, location of segment files
	tools.CreateDir(config.TEMP_DIR)

	//Dial UploadMaster
	for i := 0; i < 1000; i++ {
		log.Println("-- dialing UPLOADMASTER")
		conn, err := net.Dial("tcp", config.UPLOADMASTER_IP + config.RPC_PORT)
		if err != nil {
			log.Println("-- cannot establish connection, redial in 3 seconds : " + config.UPLOADMASTER_IP)
			time.Sleep(3* time.Second) //3 second delay
		} else {
			log.Println("-- dial finished, UPLOADMASTER is connected")
			//Create a listener to handle a call from server
			client := rpc2.NewClient(conn)


			//Register methods that client handle
			client.Handle("Alive", func(client *rpc2.Client, args *Args, reply *Reply) error {
				log.Println("-- checked from UPLOADMASTER")
				*reply = Reply{
					Type: "S",
					IsAlive: true,
				}
				return nil
			})
			client.Handle("Split", Split)

			client.Run()


		}
	}

	log.Fatal("-- Internal Network Failure, UPLOADMASTER is unresponsive")
	return

}

type Args struct{}
type Reply struct{
	Type 	string
	IsAlive	bool
}
