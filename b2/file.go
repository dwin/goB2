package b2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	t "github.com/dwin/goB2/b2/b2Types"
)

func (creds *Credential) DeleteFileVersion(filename, fileID string) (err error) {
	// Authorize and Get API Token
	err = creds.authorize()
	if err != nil {
		return err
	}

	// Create JSON body
	body := bytes.NewBuffer([]byte(`{"fileName": "` + filename + `", "fileId":"` + fileID + `"}`))

	// Create client
	client := &http.Client{}

	// Create request to (POST https://api001.backblazeb2.com/b2api/v1/b2_delete_file_version)
	req, err := http.NewRequest("POST", creds.APIAuth.APIURL+"/b2api/v1/b2_delete_file_version", body)
	fmt.Println(creds.APIAuth.APIURL + "/b2api/v1/b2_delete_file_version")
	// Headers
	req.Header.Add("Authorization", creds.APIAuth.AuthorizationToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Could not complete delete file version request. Error: %s", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return fmt.Errorf("Could not delete file version. API Resp Body: %s", string(respBody))
	}

	var file t.File

	// Parse JSON 'Bucket' Response
	err = json.Unmarshal(respBody, &file)
	if err != nil {
		return fmt.Errorf("Could not unmarshall delete file version response JSON. Err: %s", err)
	}

	return err
}
