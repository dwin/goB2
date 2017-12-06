package b2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	t "github.com/dwin/goB2/b2/b2Types"
)

// NewUploadURL requests Upload URL from API and returns 'Bucket' containing Bucket.Auth
func (creds *Credential) NewUploadURL(bucketID string) (bucket t.Bucket, err error) {
	bucket.BucketID = bucketID
	bucket.AccountID = creds.AcctID

	// Authorize and Get API Token
	err = creds.authorize()
	if err != nil {
		return bucket, err
	}

	// Get Upload URL (POST https://api001.backblazeb2.com/b2api/v1/b2_get_upload_url)
	// Create JSON body
	body := bytes.NewBuffer([]byte(`{"bucketId": "` + bucketID + `"}`))
	// Create client
	client := &http.Client{}
	// Create request
	req, err := http.NewRequest("POST", creds.APIAuth.APIURL+"/b2api/v1/b2_get_upload_url", body)

	// Headers
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return bucket, fmt.Errorf("NewUploadURL request failure. Err: %s", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(respBody, &bucket.Upload)
	if err != nil {
		return bucket, err
	}

	return bucket, err
}

// NewUploadPartURL requests new part upload url from API and returns 'UploadAuth'
func (creds *Credential) NewUploadPartURL(fileID string) (partAuth t.UploadAuth, err error) {
	// Authorize and Get API Token
	err = creds.authorize()
	if err != nil {
		return partAuth, err
	}

	// Get Upload URL (POST https://api001.backblazeb2.com/b2api/v1/b2_get_upload_part_url)
	// Create JSON body
	body := bytes.NewBuffer([]byte(`{"fileId": "` + fileID + `"}`))
	// Create client
	client := &http.Client{}
	// Create request
	req, err := http.NewRequest("POST", creds.APIAuth.APIURL+"/b2api/v1/b2_get_upload_part_url", body)

	// Headers
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return partAuth, fmt.Errorf("NewUploadPartURL request failure. Err: %s", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(respBody, &partAuth)
	if err != nil {
		return partAuth, err
	}

	return partAuth, err
}
