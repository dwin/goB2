package b2

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// DownloadByID retrieves file from B2 and saves to destination directory, if destinationFilename string is empty the filename on B2 will be used
func (creds *Credential) DownloadByID(fileID, destinationDir, destinationFilename string) (err error) {
	// Verify desination is Directory
	info, err := os.Lstat(destinationDir)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("Desination Directory given is not a directory")
	}

	// Authorize and Get API Token
	err = creds.authorize()
	if err != nil {
		return err
	}

	// Create JSON body
	body := bytes.NewBuffer([]byte(`{"fileId": "` + fileID + `"}`))

	// Create client
	client := &http.Client{}

	// Create request to (POST https://f001.backblazeb2.com/b2api/v1/b2_download_file_by_id)
	req, err := http.NewRequest("POST", creds.APIAuth.DownloadURL+"/b2api/v1/b2_download_file_by_id", body)
	fmt.Println(creds.APIAuth.DownloadURL + "/b2api/v1/b2_download_file_by_id")
	// Headers
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Could not complete download file by ID request. Error: %s", err)
	}

	// Read Response Body, Save to Destination
	if resp.Status != "200 OK" {
		respBody, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return fmt.Errorf("Could not retrieve file by ID. API Resp Body: %s", string(respBody))
	}

	// Use B2 filename or destinationFilename
	var filename string
	if len(destinationFilename) > 0 {
		filename = destinationFilename
	} else {
		filename = resp.Header.Get("X-Bz-File-Name")
	}
	file, err := os.OpenFile(destinationDir+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return err
}
