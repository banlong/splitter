package innerhttp

import (
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io"
	"splitter/config"
	"splitter/context"
)

//THIS MODULE CONTAINS METHODS FOR HTTP UPLOAD TO THE DATA PROVIDER
// Upload a file with BIN, this only works with small file size and small amount of file
// Otherwise it will overuse of RAM
func PostFile(args *context.UploadArg, targetUrl string) error {
	//log.Println("-- post: ", args.BucketId)
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	targetUrl += "/upload/" + args.VideoId

	// Write field values
	bodyWriter.WriteField("segmentId", args.SegmentId)
	bodyWriter.WriteField("bucketId", args.BucketId)
	bodyWriter.WriteField("langId", args.LangId)
	bodyWriter.WriteField("repId", args.RepId)
	bodyWriter.WriteField("videoId", args.VideoId)

	var uploadFile string
	var fileId string
	switch args.BucketId{
	case config.MP4_BUCKET:
		uploadFile = args.Video
		fileId = "video"
	case config.PL_BUCKET:
		uploadFile = args.Playlist
		fileId = "playlist"
	case config.THUMBNAIL_BUCKET:
		uploadFile = args.Thumbnail
		fileId = "thumbnail"
	case config.TS_BUCKET:
		uploadFile = args.Video
		fileId = "video"
	case config.MPDV_BUCKET:
		uploadFile = args.Video
		fileId = "video"
	case config.MPDA_BUCKET:
		uploadFile = args.Audio
		fileId = "audio"
	case config.MPD_BUCKET:
		uploadFile = args.Playlist
		fileId = "playlist"
	}


	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile(fileId, uploadFile)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(uploadFile)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()


	resp, err := http.Post(targetUrl, contentType, bodyBuf)


	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("-- upload %s completed", uploadFile)
	return nil
}

// Upload a file with File location only,
// This version will work when using share cluster drive, no big data travel
// Otherwise it will overuse of RAM
func PostFileCluster(args *context.UploadArg, targetUrl string) error {
	//log.Println("-- post: ", args.BucketId)
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	targetUrl += "/uploadcluster/" + args.VideoId

	// Write field values
	bodyWriter.WriteField("segmentId", args.SegmentId)
	bodyWriter.WriteField("bucketId", args.BucketId)
	bodyWriter.WriteField("langId", args.LangId)
	bodyWriter.WriteField("repId", args.RepId)
	bodyWriter.WriteField("videoId", args.VideoId)

	var uploadFile string
	var fileId string
	switch args.BucketId{
	case config.MP4_BUCKET:
		uploadFile = args.Video
		fileId = "video"
	case config.PL_BUCKET:
		uploadFile = args.Playlist
		fileId = "playlist"
	case config.THUMBNAIL_BUCKET:
		uploadFile = args.Thumbnail
		fileId = "thumbnail"
	case config.TS_BUCKET:
		uploadFile = args.Video
		fileId = "video"
	case config.MPDV_BUCKET:
		uploadFile = args.Video
		fileId = "video"
	case config.MPDA_BUCKET:
		uploadFile = args.Audio
		fileId = "audio"
	case config.MPD_BUCKET:
		uploadFile = args.Playlist
		fileId = "playlist"
	}

	bodyWriter.WriteField(fileId, uploadFile)
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()


	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//log.Printf("-- upload %s completed", uploadFile)
	return nil
}

// Get a file
func GetFile(targetUrl string) ([]byte, error) {
	response, err := http.Get(targetUrl)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return contents, nil
	}
}

