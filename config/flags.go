package config
import (
	"log"
	"os"
)

func init() {
	initNetConfig()
	//Change setup here to have different configuration
	CONFIG = Architect{
		SELF: 			ON_LOCAL,
		UPLOADMASTER: 		ON_LOCAL,
		DB_SERVER: 		ON_LOCAL,
		DB_API: 		"RPC",
	}


	//Get config base on set up above
	UPLOADMASTER_IP = UPLOADER_LOC[CONFIG.UPLOADMASTER]
	DATASERVER_DNS 	= DATASERVER_HTTP_ADDR[CONFIG.DB_SERVER]
	DATASERVER_IP 	= DATASERVER_IPS[CONFIG.DB_SERVER]
	FFMPEG 		= FFMPEG_LOC[CONFIG.SELF]
	FFPROBE 	= FFPROBE_LOC[CONFIG.SELF]
	THEORA 		= FFTHEORA_LOC[CONFIG.SELF]
	STREAM_URL 	= STREAM_URLS[CONFIG.SELF]

	//Setup temporary directory for FFMPEG
	TEMP_DIR = TEMP_DIRS[CONFIG.SELF]
	if(TEMP_DIR == ""){
		log.Println("Cluster Drive is not set. Please add an environment 'TEMP_DIR'")
		os.Exit(1)
	}
}

