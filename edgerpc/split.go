package edgerpc
import (
	"splitter/lib/cenkalti/rpc2"
	"splitter/context"
	"splitter/config"
	"log"
	"splitter/encode"
	"errors"
)


// This RPC is called by Upload master WEB APP--> Handle all upload requests
func Split(client *rpc2.Client, args *context.SplitArgs, reply *context.SplitReply)(err error){

	log.Println("-->Split request arrived")

	//Create segment dir
	videoId := args.VideoId
	videoDir := config.TEMP_DIR + videoId + "/"
	segmentDir := videoDir + config.SEG_DIR


	//Split
	names, err := ffmpeg.Split(args.VideoFile, config.DURATION, segmentDir, videoId, ".mp4")
	if(err != nil){
		return err
	}

	//Create m3u8 & save to Bolt
	err = ffmpeg.CreateM3U8(segmentDir, videoDir, videoId)
	if(err != nil){
		return err
	}

	//Create thumbnail
	_, err = ffmpeg.GenerateThumbnail(args.VideoFile, videoDir, videoId)
	if err != nil {
		return errors.New("FFMPEG Create-Thumbnail error:" + err.Error())
	}

	//Save M3u8
	playlistPath := videoDir + videoId + ".m3u8"
	thumbnailPath := videoDir + videoId + ".png"
	err = SaveM3U8(playlistPath, videoId)
	if(err != nil){
		return err
	}

	//Save Thumbnail
	err = SaveThumbnail(thumbnailPath, videoId)
	if(err != nil){
		return err
	}



	//Generate meta data
	meta, err := ffmpeg.GenerateMediaInfo(args.VideoFile, videoId)
	if (err != nil){
		return err
	}
	meta.VideoId = args.VideoId
	meta.LibraryId = args.LibraryId
	meta.DomainId = args.DomainId
	meta.TitleShort = args.TitleShort
	meta.UserId = args.UserId

	//Create response
	segCount := names.Len()
	ret := context.SplitReply{
		From:		args.From,
		VideoFile: 	args.VideoFile,
		SegmentCount:	segCount,
		SegmentLocation:segmentDir,
		Meta: 		meta,
	}

	*reply = ret
	return nil
}
