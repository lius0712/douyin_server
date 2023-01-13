package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/lius0712/douyin_server/cmd/feed/dal/db"
	publishDb "github.com/lius0712/douyin_server/cmd/publish/dal/db"
	"github.com/lius0712/douyin_server/kitex_gen/publish"
	"github.com/lius0712/douyin_server/pkg/tencent"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"os"
	"strings"
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
	url := tencent.GetPreSignedURL(fileName, 0)
	playUrl := strings.Split(url.String(), "?")[0]
	uuidJpeg, err := uuid.NewV4()
	if err != nil {
		return err
	}
	//获取封面
	coverName := uuidJpeg.String() + ".jpg"
	coverData, err := ExtractFrame(url.String())
	if err != nil {
		return err
	}
	coverReader := bytes.NewReader(coverData)
	//上传封面
	err = tencent.UploadFile(coverName, coverReader)
	if err != nil {
		return err
	}
	cUrl := tencent.GetPreSignedURL(coverName, 0)
	coverUrl := strings.Split(cUrl.String(), "?")[0]
	video := &db.Video{
		AuthorId:      uid,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}
	return publishDb.CreateVideo(s.ctx, video)
}

// ExtractFrame 提取videoPath视频的第n帧保存至framePath中
func ExtractFrame(videoPath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg_go.Input(videoPath).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg_go.KwArgs{
			"vframes": "1",
			"format":  "image2", "vcodec": "mjpeg",
		}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	return buf.Bytes(), err
}
