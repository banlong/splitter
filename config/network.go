package config

const (
	//-----MEDIA CONFIGURATION -----//
	MP4_BOX 	string = "MP4Box"

	//----------TEMP DIRECTORY FOR TRANSCODING -------//
	HSL_DIR    string = "hsl/"
	MP4_DIR    string = "mp4/"
	SEG_DIR	   string = "segments/"
	OUT_DIR    string = "whole/"
	NAMES_FILE string = "list.txt"
	DURATION   int    = 3 //seconds

	//-------DATABASE CONFIGURATION ----- //
	//LOCAL_DB_DIR = "/mnt/mediacluster/database/"
	//PRO_DB_DIR   = "/mnt/mediacluster/database/"

	//---------MEDIA BUCKETS -----------------//
	MP4_BUCKET       	= "mp4"
	PL_BUCKET        	= "m3u8"
	TS_BUCKET        	= "segment"
	MPD_BUCKET        	= "mpegdash"
	MPDV_BUCKET        	= "mpegdashVideo"
	MPDA_BUCKET        	= "mpegdashAudio"
	THUMBNAIL_BUCKET 	= "thumbnail"

	//-----------UPLOAD METHODS-----------------//
	MP4_RPC		string	= "mp4-rpc"
	SEG_RPC		string	= "segment-rpc"
	MP4_HTTP	string	= "mp4-http"
	SEG_HTTP	string	= "segment-http"
	LIVE		string	= "live-streaming"

	//-------------Spew constant---------------------//
	CLEAN     bool = false //indicates a clean (no db file) start-up
	USERv     bool = false
	LOGINv    bool = false
	GROUPSv   bool = false
	TRACEv    bool = true
	REGISTERv bool = false
	CHANNELv  bool = false
	VIDEOv    bool = false
	STATIONv  bool = false
	FILEv     bool = false
	INDEXv    bool = false
	COMMENTSv bool = false

	//---------SERVER LOCATIONS-----------------//
	ON_LOCAL 			= "LOCAL"
	ON_LAN 				= "LAN"
	ON_BETA          	= "ON_BETA"
	ON_LAB				= "ON_LAB"
	ON_WWW				= "ON_WWW"
	ON_A             	= "ON_A"
	ON_B                = "ON_B"

	//------------SERVICE PORTS----------------//
	RPC_PORT = ":9595"
	DBSERVER_RPC_PORT = ":9098"
	DBSERVER_HTTP_PORT = ":9097"

	MAX_THREAD = 20
)
//----INTERNAL IDENTIFIER----//
var (
	UPLOADER_LOC 		map[string]string
	DATASERVER_HTTP_ADDR 	map[string] string
	DATASERVER_IPS 		map[string]string
	FFMPEG_LOC              map[string]string
	FFPROBE_LOC             map[string]string
	FFTHEORA_LOC            map[string]string
	TEMP_DIRS               map[string]string
	STREAM_URLS		map[string]string

	CONFIG			Architect
	UPLOADMASTER_IP 	string
	DATASERVER_DNS 		string
	DATASERVER_IP 		string

	FFMPEG     		string
	FFPROBE    		string
	TEMP_DIR   		string
	THEORA     		string
	STREAM_URL 		string
)

func initNetConfig() {
	UPLOADER_LOC = make(map[string]string)
	UPLOADER_LOC[ON_BETA] = "beta.cut2it.com"
	UPLOADER_LOC[ON_LAB] = "10.0.0.60"
	UPLOADER_LOC[ON_WWW] = "54.208.178.138"
	UPLOADER_LOC[ON_LOCAL] = "localhost"
	UPLOADER_LOC[ON_LAN] = "192.168.10.7"

	DATASERVER_HTTP_ADDR = make(map[string]string)
	DATASERVER_HTTP_ADDR[ON_BETA] = "http://beta.cut2it.com"
	DATASERVER_HTTP_ADDR[ON_LAB] = "http://cdn.cut2it.com"
	DATASERVER_HTTP_ADDR[ON_WWW] = "http://cut2it.com"
	DATASERVER_HTTP_ADDR[ON_LOCAL] = "http://localhost"
	DATASERVER_HTTP_ADDR[ON_LAN] = "http://192.168.10.7"

	DATASERVER_IPS = make(map[string]string)
	DATASERVER_IPS[ON_BETA] = "beta.cut2it.com"
	DATASERVER_IPS[ON_LAB] = "cdn.cut2it.com"
	DATASERVER_IPS[ON_WWW] = "cut2it.com"
	DATASERVER_IPS[ON_LOCAL] = "localhost"
	DATASERVER_IPS[ON_LAN] = "192.168.10.7"

	FFMPEG_LOC = make(map[string]string)
	FFMPEG_LOC[ON_BETA] = "/opt/ffmpeg/ffmpeg"
	FFMPEG_LOC[ON_LAB] = "/opt/ffmpeg/ffmpeg"
	FFMPEG_LOC[ON_WWW] = "/opt/ffmpeg/ffmpeg"
	FFMPEG_LOC[ON_LOCAL] = "ffmpeg"
	FFMPEG_LOC[ON_LAN] = "ffmpeg"


	FFPROBE_LOC = make(map[string]string)
	FFPROBE_LOC[ON_BETA] = "/opt/ffmpeg/ffprobe"
	FFPROBE_LOC[ON_LAB] = "/opt/ffmpeg/ffprobe"
	FFPROBE_LOC[ON_WWW] = "/opt/ffmpeg/ffprobe"
	FFPROBE_LOC[ON_LOCAL] = "ffprobe"
	FFPROBE_LOC[ON_LAN] = "ffprobe"

	FFTHEORA_LOC = make(map[string]string)
	FFTHEORA_LOC[ON_BETA] = "/opt/ffmpeg/ffmpeg2theora"
	FFTHEORA_LOC[ON_LAB] = "/opt/ffmpeg/ffmpeg2theora"
	FFTHEORA_LOC[ON_WWW] = "/opt/ffmpeg/ffmpeg2theora"
	FFTHEORA_LOC[ON_LOCAL] = "ffmpeg2theora"
	FFTHEORA_LOC[ON_LAN] = "ffmpeg2theora"

	TEMP_DIRS = make(map[string]string)
	TEMP_DIRS[ON_BETA] = "/mnt/mediacluster/temp/"
	TEMP_DIRS[ON_LAB] = "/mnt/mediacluster/temp/"
	TEMP_DIRS[ON_WWW] = "/mnt/mediacluster/temp/"
	TEMP_DIRS[ON_LOCAL] =  "/mnt/mediacluster/temp/"
	TEMP_DIRS[ON_LAN] =  "/mnt/mediacluster/temp/"

	STREAM_URLS = make(map[string]string)
	STREAM_URLS[ON_BETA] = "http://beta.cut2it.com:9097"
	STREAM_URLS[ON_LAB] = "http://beta.cut2it.com:9097"
	STREAM_URLS[ON_WWW] = "http://beta.cut2it.com:9097"
	STREAM_URLS[ON_LOCAL] = "http://localhost:9097"
	STREAM_URLS[ON_LAN] = "http://192.168.10.16:9097"
}

type Architect struct{
	SELF 					string
	UPLOADMASTER 			string
	DB_SERVER				string
	DB_API					string
}
