package b2Types

// APIAuthorization ... Minimum Part Size deprecated and will match recommended part size
type APIAuthorization struct {
	AccountID           string `json:"accountId"`
	APIURL              string `json:"apiUrl"`
	AuthorizationToken  string `json:"authorizationToken"`
	DownloadURL         string `json:"downloadURL"`
	MinimumPartSize     int    `json:"minimumPartSize"`
	RecommendedPartSize int    `json:"recommendedPartSize"`
	AbsoluteMinPartSize int    `json:"absoluteMinimumPartSize"`
}
type UploadAuth struct {
	AuthorizationToken string `json:"authorizationToken"`
	BucketId           string `json:"bucketId"`
	URL                string `json:"uploadUrl"`
}
