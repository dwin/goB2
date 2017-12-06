package b2

import (
	"testing"
)

func TestToDeleteFileVersion(t *testing.T) {
	cred, err := New("", "", "")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	err = cred.DeleteFileVersion("The Subtle Art of Not Giving a F ck by Mark Manson.m4b", "4_zb6ee61624837a6c6588b0715_f203a13f3fbc9214d_d20170831_m005510_c001_v0001090_t0027")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
