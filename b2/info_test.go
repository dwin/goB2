package b2

import (
	"testing"
)

func TestToGetFiles(t *testing.T) {
	cred, err := New("", "", "")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	files, err := cred.GetFiles("b6ee61624837a6c6588b0715", "")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(files)
	PrintFiles(files)
}

func TestCreateBucket(t *testing.T) {
	cred, err := New("", "", "")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	bucket, err := cred.CreateBucket("testbucket", true)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(bucket.BucketID)
}

func TestGetBuckets(t *testing.T) {
	cred, err := New("", "", "")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	buckets, err := cred.GetBuckets()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(buckets)
	PrintBuckets(buckets)
}
