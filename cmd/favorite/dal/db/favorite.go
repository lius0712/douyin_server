package db

import (
	"context"
	dbVideo "github.com/lius0712/douyin_server/cmd/feed/dal/db"
	dbUser "github.com/lius0712/douyin_server/cmd/user/dal/db"
	"github.com/lius0712/douyin_server/pkg/errno"
	"gorm.io/gorm"
)

func Favorite(ctx context.Context, uid int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		user := new(dbUser.User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}
		video := new(dbVideo.Video)
		if err := tx.WithContext(ctx).First(video, vid).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&user).Association("FavoriteVideos").Append(video); err != nil {
			return err
		}
		//修改video表中的favorite count
		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}
		return nil
	})
	return err
}
func DisFavorite(ctx context.Context, uid int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		user := new(dbUser.User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}
		//TODO: 修改 video 表中的 favorite count
		return nil
	})
	return err
}
func FavoriteList(ctx context.Context, uid int64) ([]dbVideo.Video, error) {
	user := new(dbUser.User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}
	videos := []dbVideo.Video{}
	if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&videos); err != nil {
		return nil, err
	}
	return videos, nil
}
