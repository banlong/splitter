package innerrpc
import "sync"

type (
	RPC struct {
		requests *Requests
		mu       *sync.RWMutex
	}

	// Monitor service activities
	Requests struct {
		Get    uint64
		Put    uint64
		Delete uint64
		Clear  uint64
	}

	GetArgs struct {
		Database 			string
		BucketId 			string
		Key                 string
	}

	GetReply struct {
		Database 			string
		BucketId 			string
		Key                 string
		Data				[]byte
	}

	//Post data to DB
	PostArgs struct {
		Database 			string
		BucketId 			string
		Key                 string
		Data				[]byte
	}

	ClusterPostArgs struct {
		Database 			string
		BucketId 			string
		Key                 string
		Data				string
	}
)


//type Args struct{
//	Method 		string `json: "method"`
//	VideoId 	string `json: "videoId"`
//	SegmentId 	string `json: "segmentId"`
//	LangId 		string `json: "langId"`
//	RepId 		string `json: "repId"`
//	Video 		[]byte
//	Audio 		[]byte
//	Thumbnail 	[]byte
//	Playlist 	[]byte
//	IsLastSegment bool   `json: "isLastSegment"`
//	LibraryId     string `json: "libraryId"`
//	UserId        string `json:"userid"`
//	DomainId      string `json:"domain"`
//	ThumbnailUrl  string `json: "thumbnailUrl"`
//	TitleShort    string `json: "titleShort"`
//	Test          bool
//}