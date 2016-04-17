package context


type VideoArgs struct {
	LibraryId        string
	VideoId          string `json:"videoid"`
	UserId           string `json:"userid"`
	DomainId         string `json:"domain"`
	AudioCodec       string `json:"acodec"`
	VideoCodec       string `json:"vcodec"`
	Orientation      string `json: "Orientation"`
	ThumbnailUrl     string `json: "ThumbnailUrl"`
	UrlMp4		     string `json: "Url"`
	Duration         string `json: "Duration"`
	TitleShort       string `json: "TitleShort"`
	Width            int    `json: "width"`
	Height           int    `json: "height"`
	TranscodePending bool
	UrlM3U8          string
}


//From UploadMaster To Splitter
type SplitArgs struct {
	From 		string `json: "from"`
	VideoId 	string `json: "videoId"`
	VideoFile	string
	LibraryId     	string `json: "libraryId"`
	UserId        	string `json:"userid"`
	DomainId      	string `json:"domain"`
	TitleShort    	string `json: "titleShort"`

}


type SplitReply struct {
	From            string `json: "from"`
	VideoFile       string
	SegmentCount    int
	SegmentLocation string
	Meta 		*VideoArgs
}



//From Splitter to Bolt
type UploadArg struct{
	VideoId 	string  `json: "videoId"`
	SegmentId	string  `json: "segmentId"`
	BucketId	string  `json: "bucketId"`
	RepId 		string  `json: "repId"`
	LangId		string  `json: "langId"`
	Video       string  `json: "video"`
	Audio 		string  `json: "audio"`
	Playlist    string  `json: "playlist"`
	Thumbnail 	string  `json: "thumbnail"`
}