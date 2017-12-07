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

	"github.com/dustin/go-humanize"
)

// GetFiles returns specified number of files in bucket starting with given name, prefix, and delimiter. BucketID is required, other variables are optional per B2 API docs.
func (creds *Credential) GetFiles(bucketID, startFileName, prefix, delimiter string, maxFileCount int) (files t.Files, err error) {
	// Authorize and Get API Token
	err = creds.authorize()
	if err != nil {
		return files, err
	}

	// Create json body          **Max File Count not working, removed temporarily
	reqjson, err := json.Marshal(t.RequestFile{
		BucketID:      bucketID,
		StartFileName: startFileName,
		//MaxFileCount:  maxFileCount,
		Prefix:    prefix,
		Delimiter: delimiter,
	})

	// Create client
	client := &http.Client{}
	// Create request
	req, err := http.NewRequest("POST", creds.APIAuth.APIURL+"/b2api/v1/b2_list_file_names", bytes.NewBuffer(reqjson))
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return files, err
	}
	// Read Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return files, fmt.Errorf("Error response from API. Err: %s", respBody)
	}
	err = json.Unmarshal(respBody, &files)
	if err != nil {
		return files, fmt.Errorf("Error parsing JSON response for request all filenames for bucket %s. Err: %s", bucketID, err)
	}

	return files, err
}

// GetFilesVersions returns specified number of files in bucket starting with given name, id, prefix, and delimiter. BucketID is required, other variables are optional per B2 API docs.
func (creds *Credential) GetFilesVersions(bucketID, startFileName, startFileID, prefix, delimiter string, maxFileCount int) (files t.Files, nextFile t.File, err error) {
	// Authorize and Get API Token
	err = creds.authorize()
	if err != nil {
		return files, nextFile, err
	}

	// Create json body          **Max File Count not working, removed: "maxFileCount":"1000",
	reqjson, err := json.Marshal(t.RequestFile{
		BucketID:      bucketID,
		StartFileName: startFileName,
		StartFileID:   startFileID,
		//MaxFileCount:  maxFileCount,
		Prefix:    prefix,
		Delimiter: delimiter,
	})

	// Create client
	client := &http.Client{}
	// Create request
	req, err := http.NewRequest("POST", creds.APIAuth.APIURL+"/b2api/v1/b2_list_file_versions", bytes.NewBuffer(reqjson))
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return files, nextFile, err
	}
	// Read Response
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return files, nextFile, fmt.Errorf("Error response from API. Err: %s", respBody)
	}
	err = json.Unmarshal(respBody, &files)
	if err != nil {
		return files, nextFile, fmt.Errorf("Error parsing JSON response for request all file versions for bucket %s. Err: %s", bucketID, err)
	}

	return files, nextFile, err
}

// CreateBucket makes new B2 bucket and returns API response and error
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

// DeleteBucket destroys B2 bucket and returns API response and error
func (creds *Credential) DeleteBucket(bucketID string) (bucket t.Bucket, err error) {
	// Validate Input
	if len(bucketID) < 1 {
		return bucket, fmt.Errorf("Bucket ID given has invalid length, too short")
	}

	// Authorize and Get API Token
	err = creds.authorize()
	if err != nil {
		return bucket, err
	}

	// Create JSON body
	body := bytes.NewBuffer([]byte(`{"accountId": "` + creds.APIAuth.AccountID + `", "bucketId":"` + bucketID + `" }`))

	// Create client
	client := &http.Client{}

	// Create request to (POST https://api001.backblazeb2.com/b2api/v1/b2_delete_bucket)
	req, err := http.NewRequest("POST", creds.APIAuth.APIURL+"/b2api/v1/b2_delete_bucket", body)
	fmt.Println(creds.APIAuth.APIURL + "/b2api/v1/b2_delete_bucket")
	// Headers
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return bucket, fmt.Errorf("Could not complete delete bucket request. Error: %s", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return bucket, fmt.Errorf("Could not delete Bucket. API Resp Body: %s", string(respBody))
	}
	// Parse JSON 'Bucket' Response
	err = json.Unmarshal(respBody, &bucket)
	if err != nil {
		return bucket, fmt.Errorf("Could not unmarshall delete bucket response JSON. Err: %s", err)
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

// GetFileInfo calls the API using the supplied File ID and returns a type 'File' and error
func (creds *Credential) GetFileInfo(fileID string) (file t.File, err error) {
	// Authorize and get API token
	err = creds.authorize()
	if err != nil {
		return file, err
	}

	// Create JSON body
	body := bytes.NewBuffer([]byte(`{"fileId": "` + fileID + `"}`))

	// Create client
	client := &http.Client{}

	// Create request to (POST https://api001.backblazeb2.com/b2api/v1/b2_get_file_info)
	req, err := http.NewRequest("POST", creds.APIAuth.APIURL+"/b2api/v1/b2_get_file_info", body)

	// Headers
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return file, fmt.Errorf("Could not complete GetFileInfo request. Err: %s", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return file, fmt.Errorf("Error response from API to GetBuckets request. API Resp: %s", string(respBody))
	}
	// Parse JSON 'Buckets' Response
	err = json.Unmarshal(respBody, &file)
	if err != nil {
		return file, fmt.Errorf("Unable to unmarshall JSON response. API Resp: %s", string(respBody))
	}

	return file, err
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

// PrintFiles Diplays list of files in console
func PrintFiles(files t.Files) {
	if len(files.File) > 1 {
		writer := new(tabwriter.Writer)
		fmt.Println("B2 Files")
		// Format to '|' separated columns with no min width and blank padding char
		writer.Init(os.Stdout, 0, 5, 1, ' ', 0)
		fmt.Fprintln(writer, "-NAME-\t -SIZE-\t -ID-\t -UPLOAD TIME-\t")
		for i := 0; i < len(files.File); i++ {
			//uploadTime := fmt.Sprintf("%s", humanize.Time(time.Unix(files.File[i].UploadTimestamp, 0)))
			// Convert MS to NS
			uploadTime := time.Unix(0, files.File[i].UploadTimestamp*1000*1000).Format(time.RFC822)
			// Display folder
			size := ""
			if files.File[i].ContentLength == 0 {
				size = "folder"
			} else {
				size = fmt.Sprintf("%s", humanize.Bytes(uint64(files.File[i].ContentLength)))
			}
			fmt.Fprintln(writer, files.File[i].FileName+"\t", size+"\t", files.File[i].FileID+"\t", uploadTime+"\t")
		}
		fmt.Fprintln(writer)
		writer.Flush()
	} else {
		fmt.Println("No files")
	}
}

// PrintAPIAuth Display API Authorization information in console
func PrintAPIAuth(auth t.APIAuthorization) {
	fmt.Println("--Backblaze B2 API Authorization--")
	fmt.Println("AccountID:\t" + auth.AccountID)
	fmt.Println("API URL:\t" + auth.APIURL)
	fmt.Println("Auth Token:\t" + auth.AuthorizationToken)
	fmt.Println("Download URL:\t" + auth.DownloadURL)
	fmt.Printf("Min. Part Size:\t %v, %v\n", auth.MinimumPartSize, humanize.Bytes(uint64(auth.MinimumPartSize)))
	fmt.Printf("Rec. Min Part Size:\t %v, %v\n", auth.RecommendedPartSize, humanize.Bytes(uint64(auth.RecommendedPartSize)))
	fmt.Printf("Absolute Min Part Size:\t %v, %v\n", auth.AbsoluteMinPartSize, humanize.Bytes(uint64(auth.AbsoluteMinPartSize)))
}
