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
	f := strings.NewReader("test11111")
	UploadFile("test", f)
}

func TestGetFile(t *testing.T) {
	GetFileUrl("test")
}
