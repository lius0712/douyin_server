package tencent

import (
	"context"
	"io"
	"net/url"
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

	url := tencentClient.Object.GetObjectURL(objectName)

	return url
}
