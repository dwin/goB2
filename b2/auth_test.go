package b2

import (
	"fmt"
	"testing"
)

func TestNewCredandAuth(t *testing.T) {
	//fmt.Println(os.Getenv("B2AcctID"))
	cred, err := New("", "", "")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	err = cred.authorize()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	t.Log("Authorize Credentials success")
	fmt.Println(cred)

}
