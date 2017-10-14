package b2Types

type Buckets struct {
	Bucket []Bucket `json:"buckets"`
}
type Bucket struct {
	AccountID      string   `json:"accountId"`
	BucketID       string   `json:"bucketId"`
	BucketName     string   `json:"bucketName"`
	BucketType     string   `json:"bucketType"`
	LifecycleRules []string `json:"lifecycleRules"`
	Revision       int      `json:"revision"`
	Upload         UploadAuth
}
