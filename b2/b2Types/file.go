package b2Types

type B2File struct {
	AccountID   string `json:"accountId"`
	BucketID    string `json:"bucketId"`
	ContentType string `json:"contentType"`
	FileID      string `json:"fileId"`
	FileInfo    struct {
		LargeFileSHA1          string `json:"large_file_sha1"`
		LastModificationMillis int64  `json:"src_last_modified_millis,string"`
	} `json:"fileInfo"`
	FileName        string `json:"fileName"`
	UploadTimestamp int64  `json:"uploadTimestamp"`
}

type AllFiles struct {
	File         []File `json:"files"`
	NextFileName string `json:"nextFileName"`
}
type File struct {
	Action        string `json:"action"`
	ContentLength int    `json:"contentLength"`
	ContentSha1   string `json:"contentSha1"`
	ContentType   string `json:"contentType"`
	FileID        string `json:"fileId"`
	FileInfo      struct {
		LargeFileSHA1         string `json:"large_file_sha1"`
		SrcLastModifiedMillis string `json:"src_last_modified_millis"`
	} `json:"fileInfo"`
	FileName        string `json:"fileName"`
	Size            int    `json:"size"`
	UploadTimestamp int64  `json:"uploadTimestamp"`
}
