package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/lius0712/douyin_server/kitex_gen/publish"
	"github.com/lius0712/douyin_server/pkg/tencent"
)

type PublishActionService struct {
	ctx context.Context
}

func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{
		ctx: ctx,
	}
}

func (s *PublishActionService) PublishAction(req *publish.PublishActionRequest, uid int) error {
	videoData := []byte(req.File)
	reader := bytes.NewReader(videoData)
	uuidFile, err := uuid.NewV4()
	if err != nil {
		return err
	}
	fileName := uuidFile.String() + ".mp4"
	//上传视频
	err = tencent.UploadFile(fileName, reader)
	if err != nil {
		return err
	}
	//获取视频链接
	playUrl := tencent.GetFileUrl(fileName)
	fmt.Println("*******")
	fmt.Println(playUrl)
	return nil
}
