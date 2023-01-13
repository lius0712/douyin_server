package tencent

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	CreateBucket()
}

func TestUploadLocalFile(t *testing.T) {
	err := UploadLocalFile("bear.mp4", "./bear.mp4")
	fmt.Println(err)
}

func TestUploadFile(t *testing.T) {
	f := strings.NewReader("testUpload")
	UploadFile("test123", f)
}

func TestGetFile(t *testing.T) {
	GetFileUrl("test")
}

func TestGetPreSignedURL(t *testing.T) {
	url := GetPreSignedURL("0ef40421-e920-4deb-8cf1-2b64634605e9.mp4", 0)
	fmt.Println(url)
}
