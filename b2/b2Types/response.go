package b2Types

import "net/http"

type Response struct {
	Header http.Header
	Status string
	Body   []byte
}

type UploadPartResponse struct {
	AuthorizationToken string `json:"authorizationToken"`
	FileID             string `json:"fileId"`
	UploadURL          string `json:"uploadUrl"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
