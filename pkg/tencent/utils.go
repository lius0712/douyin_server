package tencent

import (
	"context"
	"github.com/lius0712/douyin_server/pkg/constants"
	"io"
	"net/http"
	"net/url"
	"time"
)

func CreateBucket() error {
	_, err := tencentClient.Bucket.Put(context.Background(), nil)
	if err != nil {
		return err
	}
	return nil
}

// 通过本地文件上传对象
func UploadLocalFile(objectName string, filePath string) error {
	ctx := context.Background()
	//filePath := "../public"
	_, err := tencentClient.Object.PutFromFile(ctx, objectName, filePath, nil)

	if err != nil {
		return err
	}
	return nil
}

// 通过文件流上传对象
func UploadFile(objectName string, f io.Reader) error {
	ctx := context.Background()
	_, err := tencentClient.Object.Put(ctx, objectName, f, nil)
	if err != nil {
		return err
	}
	return nil
}

func GetFileUrl(objectName string) *url.URL {

	//opt := &cos.BucketGetOptions{
	//	Prefix: objectName,
	//}
	//v, _, err := tencentClient.Bucket.Get(context.Background(), opt)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(v.Name)
	//fmt.Println(v.Prefix)
	//
	//for _, c := range v.Contents {
	//	fmt.Printf("%s, %s, %d\n", c.Key, c.StorageClass, c.Size)
	//}
	playUrl := tencentClient.Object.GetObjectURL(objectName)
	return playUrl
}

func GetPreSignedURL(objectName string, expires time.Duration) *url.URL {
	ctx := context.Background()
	ak := constants.SecretID
	sk := constants.SecretKey
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	presignedURL, err := tencentClient.Object.GetPresignedURL(ctx, http.MethodGet, objectName, ak, sk, expires, nil)
	if err != nil {
		panic(err)
	}
	return presignedURL
}
