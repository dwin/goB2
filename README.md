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

**Implemented API Operations:**

- [x] b2_authorize_account
- [ ] b2_cancel_large_file
- [x] b2_create_bucket
- [x] b2_delete_bucket
- [x] b2_delete_file_version
- [x] b2_download_file_by_id
- [ ] b2_download_file_by_name
- [ ] b2_get_download_authorization
- [ ] b2_finish_large_file
- [x] b2_get_file_info
- [x] b2_get_upload_part_url
- [x] b2_get_upload_url
- [ ] b2_hide_file
- [x] b2_list_buckets
- [x] b2_list_file_names
- [x] b2_list_file_versions
- [ ] b2_list_parts
- [ ] b2_list_unfinished_large_files
- [ ] b2_start_large_file
- [ ] b2_update_bucket
- [ ] b2_upload_file
- [ ] b2_upload_part

**To Do:**

- [ ] Find Bucket ID by name
- [ ] Find FileIDs for given filename
- [ ] Find Filename for FileID
- [ ] *Fix:* Using MaxFileCount variable on list files or versions returns unexpected response from API, pending B2 Technical Desk response
- [ ] Verify SHA-1 on Download