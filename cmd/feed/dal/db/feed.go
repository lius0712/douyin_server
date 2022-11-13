package db

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/user/dal/db"
	"github.com/lius0712/douyin_server/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	AuthorId      int     `gorm:"index:idx_authorId;not null"`
	Author        db.User `gorm:"foreignKey:AuthorId"`
	PlayUrl       string  `gorm:"type:varchar(255);not null"`
	CoverUrl      string  `gorm:"type:varchar(255)"`
	FavoriteCount int     `gorm:"default:0"`
	CommentCount  int     `gorm:"default:0"`
	Title         string  `gorm:"type:varchar(50);not null"`
}

func (Video) TableName() string {
	return constants.VideoTableName
}

func MGetVideos(ctx context.Context, limit int, latestTime *int64) ([]*Video, error) {
	videos := make([]*Video, 0)

	//如果latestTime为空或值为0
	if latestTime == nil || *latestTime == 0 {
		curTime := time.Now().UnixMilli()
		latestTime = &curTime
	}

	if err := DB.WithContext(ctx).Limit(limit).Order("created_at desc").Find(&videos, "created_at < ?", time.UnixMilli(*latestTime)).Error; err != nil {
		return nil, err
	}

	return videos, nil
}
