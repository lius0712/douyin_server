package tencent

import (
	"context"
	"github.com/lius0712/douyin_server/pkg/constants"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var tencentClient *cos.Client

func init() {

	u, _ := url.Parse(constants.CosUrl)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  constants.SecretID,
			SecretKey: constants.SecretKey,
		},
	})

	//检索存储桶是否存在且是否有权限访问
	_, err := client.Bucket.Head(context.Background())

	if err != nil {
		panic(err)
	}

	tencentClient = client

}
