package b2

import (
	"testing"
)

func TestToDownloadFileByID(t *testing.T) {
	cred, err := New("", "", "")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	err = cred.DownloadByID("4_zb6ee61624837a6c6588b0715_f1163aebca70b6036_d20161202_m004913_c001_v0001032_t0008", "/Users/dwin/Dev/golang/src/github.com/dwin/goB2/b2", "test.pdf")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log("File Download By ID success")
}
