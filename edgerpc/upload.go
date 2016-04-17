package edgerpc

import (
	"splitter/config"
	"splitter/context"
	"splitter/innerhttp"
	"log"
	"splitter/innerrpc"
)

func SaveM3U8(filepath string, videoId string)  error{

	if(config.CONFIG.DB_API == "HTTP"){
		playlistArg := context.UploadArg{
			VideoId: 	videoId,
			BucketId:	config.PL_BUCKET,
			Playlist:   	filepath,
		}
		err := innerhttp.PostFileCluster(&playlistArg, config.DATASERVER_DNS + config.DBSERVER_HTTP_PORT)
		if(err != nil){
			return err
		}
		log.Println("-- save ",  videoId + ".m3u8")

	}else{
		m3u8Arg := innerrpc.ClusterPostArgs{
			Database:	videoId,
			BucketId: 	config.PL_BUCKET,
			Key: 		videoId + ".m3u8",
			Data: 		filepath,
		}
		_, err :=  innerrpc.DataClient.ClusterPut(&m3u8Arg)
		if(err != nil){
			return err
		}
		log.Println("-- save ",  videoId + ".m3u8")

	}
	return  nil
}

func SaveThumbnail(filepath string, videoId string)  error{

	if(config.CONFIG.DB_API == "HTTP"){
		//Save Thumbnail
		thumbArg := context.UploadArg{
			VideoId: 	videoId,
			BucketId:	config.THUMBNAIL_BUCKET,
			Thumbnail:  	filepath,
		}
		err := innerhttp.PostFileCluster(&thumbArg, config.DATASERVER_DNS + config.DBSERVER_HTTP_PORT)
		if(err != nil){
			return err
		}
		log.Println("-- save ",  videoId + ".m3u8")

	}else{
		thumArg := innerrpc.ClusterPostArgs{
			Database:	videoId,
			BucketId: 	config.THUMBNAIL_BUCKET,
			Key: 		videoId + ".png",
			Data: 		filepath,
		}
		_, err :=  innerrpc.DataClient.ClusterPut(&thumArg)
		if(err != nil){
			return err
		}
		log.Println("-- save ",  videoId + ".png")

	}
	return  nil
}
