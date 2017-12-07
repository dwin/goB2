package b2Types

type Files struct {
	File         []File `json:"files"`
	NextFileName string `json:"nextFileName"`
}

type File struct {
	Action          string            `json:"action,omitempty"`
	AccountID       string            `json:"accountId,omitempty"`
	BucketID        string            `json:"bucketId,omitempty"`
	ContentLength   int               `json:"contentLength"`
	ContentSha1     string            `json:"contentSha1"`
	ContentType     string            `json:"contentType"`
	FileID          string            `json:"fileId"`
	FileInfo        map[string]string `json:"fileInfo,omitempty"`
	FileName        string            `json:"fileName"`
	Size            int               `json:"size,omitempty"`
	UploadTimestamp int64             `json:"uploadTimestamp"`
}

type RequestFile struct {
	BucketID      string `json:"bucketId"`
	StartFileName string `json:"startFileName"`
	StartFileID   string `json:"startFileId,omitempty"`
	MaxFileCount  int    `json:"maxFileCount,omitempty"`
	Prefix        string `json:"prefix"`
	Delimiter     string `json:"delimiter,omitempty"`
}

/*
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
*/
