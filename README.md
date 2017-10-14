# goB2

## Usage

```go
acct, err := New("", "", "")
if err != nil {
    // Handle Error
}
bucket, err := acct.CreateBucket("testbucket", true)
if err != nil {
    // Handle Error
}
// Use bucket
fmt.Println("BucketID: "+ bucket.BucketID)
```