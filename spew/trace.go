package spew

import (
	"log"
	"nemean/config"
)

func WriteHead(s string) {
	if config.TRACEv == true {
		log.Printf("---------------UPLOAD:: ", s)
	}
}
func WriteTail() {
	if config.TRACEv == true {
		log.Printf("--------------END")
	}
}
