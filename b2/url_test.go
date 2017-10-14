package b2

import (
	"fmt"
	"testing"
)

func TestNewUploadURL(t *testing.T) {
	cred, err := New("", "", "")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	err = cred.authorize()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	b, err := cred.NewUploadURL("b6ee61624837a6c6588b0715")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	fmt.Println(b.Upload.URL)
	t.Log(b)
}
