package db

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/feed/dal/db"
	"gorm.io/gorm"
)

func CreateVideo(ctx context.Context, video *db.Video) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(video).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func PublishList(ctx context.Context, authorId int64) ([]*db.Video, error) {
	var videoList []*db.Video
	err := DB.WithContext(ctx).Model(&db.Video{}).Where(&db.Video{
		AuthorId: int(authorId),
	}).Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
