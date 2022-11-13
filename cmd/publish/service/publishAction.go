package service

import (
	"context"
	"github.com/lius0712/douyin_server/kitex_gen/publish"
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
	//videoData := []byte(req.File)
	//
	//reader := bytes.NewReader(videoData)
	//
	//uuidVideo, err := uuid.NewV4()
	//
	//if err != nil {
	//	return err
	//}
	//
	//videoName := uuidVideo.String() + ".mp4"

	return nil
}
