package b2

import (
	"fmt"
	"testing"
)

func TestToGetFiles(t *testing.T) {
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
	files, err := cred.GetFiles(buckets.Bucket[0].BucketID, "")
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

func TestGetFileInfo(t *testing.T) {
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
	files, err := cred.GetFiles(buckets.Bucket[0].BucketID, "")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	file, err := cred.GetFileInfo(files.File[0].FileID)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(file)
	fmt.Println("File Name: " + file.FileName)
}
