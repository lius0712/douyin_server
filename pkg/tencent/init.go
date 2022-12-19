package tencent

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var tencentClient *cos.Client

func init() {

	u, _ := url.Parse("")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "",
			SecretKey: "",
		},
	})

	//检索存储桶是否存在且是否有权限访问
	_, err := client.Bucket.Head(context.Background())

	if err != nil {
		panic(err)
	}

	tencentClient = client

}
