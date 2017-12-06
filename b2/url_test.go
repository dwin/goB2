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
		t.FailNow()
	}
	b, err := cred.NewUploadURL("b6ee61624837a6c6588b0715")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	fmt.Println(b.Upload.URL)
	t.Log(b)
}

func TestNewUploadPartURL(t *testing.T) {
	cred, err := New("", "", "")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	err = cred.authorize()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	partAuth, err := cred.NewUploadPartURL("4_zb6ee61624837a6c6588b0715_f20317adb7708141a_d20170209_m184527_c001_v0001029_t0013")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	fmt.Println("Part Auth URL: ", partAuth.URL)
	t.Log(partAuth)
}
