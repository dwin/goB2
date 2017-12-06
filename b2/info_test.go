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

func TestDeleteBucket(t *testing.T) {
	cred, err := New("", "", "")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	bucket, err := cred.DeleteBucket("f6ae41d288a786a6589b0715")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log("Deleted", bucket.BucketID)
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
