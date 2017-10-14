package b2

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	t "github.com/dwin/goB2/b2/b2Types"
)

type Credential struct {
	AcctID  string
	AppKey  string
	APIURL  string
	APIAuth t.APIAuthorization
}

// New returns new B2 Credential
// Optionally you can specify empty strings is you would like to use environmental variables instead of B2 credentials in code
func New(acctID, appKey, apiURL string) (*Credential, error) {
	if acctID == "" {
		acctID = os.Getenv("B2AcctID")
	}
	if appKey == "" {
		appKey = os.Getenv("B2AppKey")
	}
	if apiURL == "" {
		apiURL = os.Getenv("B2APIURL")
	}
	if len(acctID) < 4 || len(appKey) < 4 || len(apiURL) < 4 {
		err := errors.New("B2 Account Credential Length Check Failed")
		return nil, err
	}
	return &Credential{
		AcctID: acctID,
		AppKey: appKey,
		APIURL: apiURL,
	}, nil
}

func (creds *Credential) authorize() error {
	// Encode credentials to base64
	b64cred := base64.StdEncoding.EncodeToString([]byte(creds.AcctID + ":" + creds.AppKey))

	// Request (POST https://api.backblazeb2.com/b2api/v1/b2_authorize_account)
	body := bytes.NewBuffer([]byte(`{}`))

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", creds.APIURL+"b2_authorize_account", body)
	if err != nil {
		return errors.New("Creating API Auth Request Failed")
	}

	// Headers
	req.Header.Add("Authorization", "Basic "+b64cred)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		return errors.New("API Auth Request Failed")
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return fmt.Errorf("Authorization with Backblaze B2 API Failed. API Resp Body: %s", string(respBody))
	}
	err = json.Unmarshal(respBody, &creds.APIAuth)
	if err != nil {
		return errors.New("Cannot parse API Auth Response JSON")
	}

	// Check API Response matches config
	if creds.APIAuth.AccountID != creds.AcctID {
		return fmt.Errorf("API Account ID Response does not match Account ID in Config. API Acct ID: %s != Config Acct ID: %s", creds.APIAuth.AccountID, creds.AcctID)
	}

	return err
}
