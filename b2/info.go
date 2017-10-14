package b2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	t "github.com/dwin/goB2/b2/b2Types"
)

// GetFiles returns all files in bucket
func (creds *Credential) GetFiles(bucketID, startFile string) ([]t.File, error) {
	// Authorize and Get API Token
	err := creds.authorize()
	if err != nil {
		return nil, err
	}

	// Create json body
	body := bytes.NewBuffer([]byte(`{"bucketId":"` + bucketID + `","startFileName":"` + startFile + `"}`))
	// Create client
	client := &http.Client{}
	// Create request
	req, err := http.NewRequest("POST", creds.APIURL+"/b2api/v1/b2_list_file_names", body)
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)
	/*
		resp, err := resty.R().
			SetHeader("Authorization", creds.APIAuth).
			SetBody(`{"bucketId":"` + bucketId + `","startFileName":"` + startFile + `"}`).
			Post(creds.APIURL + "/b2api/v1/b2_list_file_names")
		if err != nil {
			logger.Fatal("API Communication Error: Could not get filename list",
				zap.Error(err),
			)
		}
	*/
	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Read Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var allFiles t.AllFiles
	err = json.Unmarshal(respBody, &allFiles)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON response for request all filenames for bucket %s", bucketID)
	}

	// Display files
	/*
		for i := 0; i < len(allFiles.File); i++ {
			fmt.Printf("\n\nFileID: %v\nFilename: %v\nSHA1: %v\nBlake2b: %v\nSize: %v",
				allFiles.File[i].FileID, allFiles.File[i].FileName, allFiles.File[i].ContentSha1,
				allFiles.File[i].FileInfo.ContentBlake2B, allFiles.File[i].Size)
		}
	*/
	return allFiles.File, err

}

// CreateBucket makes new B2 bucket and returns API response
func (creds *Credential) CreateBucket(bucketName string, bucketPublic bool) (bucket t.Bucket, err error) {
	//TODO: Check bucket name validity

	if len(bucketName) < 6 {
		return bucket, fmt.Errorf("Bucket Name must be at least 6 chars")
	}

	// Public or private bucketName
	var bucketType = "allPrivate"
	if bucketPublic == true {
		bucketType = "allPublic"
	}

	// Authorize and Get API Token
	err = creds.authorize()
	if err != nil {
		return bucket, err
	}

	// Create JSON body
	body := bytes.NewBuffer([]byte(`{"accountId": "` + creds.APIAuth.AccountID + `", "bucketName":"` + bucketName + `", "bucketType":"` + bucketType + `" }`))

	// Create client
	client := &http.Client{}

	// Create request to (POST https://api001.backblazeb2.com/b2api/v1/b2_create_bucket)
	req, err := http.NewRequest("POST", creds.APIAuth.APIURL+"/b2api/v1/b2_create_bucket", body)
	fmt.Println(creds.APIAuth.APIURL + "/b2api/v1/b2_create_bucket")
	// Headers
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return bucket, fmt.Errorf("Could not complete create bucket request. Error: %s", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return bucket, fmt.Errorf("Could not create new Bucket. API Resp Body: %s", string(respBody))
	}
	// Parse JSON 'Bucket' Response
	err = json.Unmarshal(respBody, &bucket)
	if err != nil {
		return bucket, fmt.Errorf("Could not unmarshall create bucket response JSON. Err: %s", err)
	}

	return bucket, err
}

// GetBuckets calls authorize then connects to API to request list of all B2 buckets and information, returns type 'Buckets' and error
func (creds *Credential) GetBuckets() (buckets t.Buckets, err error) {
	// Authorize and Get API Token
	err = creds.authorize()
	if err != nil {
		return buckets, err
	}

	// Create JSON body
	body := bytes.NewBuffer([]byte(`{"accountId": "` + creds.APIAuth.AccountID + `"}`))

	// Create client
	client := &http.Client{}

	// Create request to (POST https://api001.backblazeb2.com/b2api/v1/b2_list_buckets)
	req, err := http.NewRequest("POST", creds.APIAuth.APIURL+"/b2api/v1/b2_list_buckets", body)

	// Headers
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return buckets, fmt.Errorf("Could not complete GetBuckets request. Err: %s", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return buckets, fmt.Errorf("Error response from API to GetBuckets request. API Resp: %s", string(respBody))
	}
	// Parse JSON 'Buckets' Response
	err = json.Unmarshal(respBody, &buckets)
	if err != nil {
		return buckets, fmt.Errorf("Unable to unmarshall JSON response. API Resp: %s", string(respBody))
	}

	return buckets, err
}

// PrintFiles Diplays list of files in console
func PrintFiles(files []t.File) {
	if len(files) > 1 {
		writer := new(tabwriter.Writer)
		fmt.Println("B2 Files")
		// Format to '|' separated columns with no min width and blank padding char
		writer.Init(os.Stdout, 0, 5, 1, ' ', 0)
		fmt.Fprintln(writer, "-NAME-\t -SIZE-\t -ID-\t -UPLOAD TIME-\t")
		for i := 0; i < len(files); i++ {
			uploadTime, _ := time.Parse("RFC822", string(files[i].UploadTimestamp))
			tS := fmt.Sprintf("%s", uploadTime)
			fmt.Fprintln(writer, files[i].FileName+"\t", string(files[i].Size)+"\t", files[i].FileID+"\t", tS+"\t")
		}
		fmt.Fprintln(writer)
		writer.Flush()
	} else {
		fmt.Println("No files")
	}
}

// PrintBuckets Diplays list of files in console
func PrintBuckets(buckets t.Buckets) {
	if buckets.Bucket != nil {
		writer := new(tabwriter.Writer)
		fmt.Println("B2 Buckets")
		// Format to '|' separated columns with no min width and blank padding char
		writer.Init(os.Stdout, 0, 5, 1, ' ', 0)
		fmt.Fprintln(writer, "-ID-\t -NAME-\t -TYPE-")
		for i := 0; i < len(buckets.Bucket); i++ {
			fmt.Fprintln(writer, buckets.Bucket[i].BucketID+"\t", buckets.Bucket[i].BucketName+"\t", buckets.Bucket[i].BucketType+"\t")
		}
		fmt.Fprintln(writer)
		writer.Flush()
	} else {
		fmt.Println("No buckets")
	}
}
