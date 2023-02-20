package db

import (
	"context"
	dbFeed "github.com/lius0712/douyin_server/cmd/feed/dal/db"
	dbUser "github.com/lius0712/douyin_server/cmd/user/dal/db"
	"github.com/lius0712/douyin_server/pkg/constants"
	"github.com/lius0712/douyin_server/pkg/errno"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Video   dbFeed.Video `gorm:"foreignkey:VideoID"`
	VideoID int          `gorm:"index:idx_videoid;not null"`
	User    dbUser.User  `gorm:"foreignkey:UserID"`
	UserID  int          `gorm:"index:idx_userid;not null"`
	Content string       `gorm:"type:varchar(255);not null"`
}

func (Comment) TableName() string {
	return constants.CommentTableName
}

func AddComment(ctx context.Context, comment *Comment) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error { //事务操作
		err := tx.Create(comment).Error
		if err != nil {
			return err
		}
		//改变video表中的comment count
		res := tx.Model(&dbFeed.Video{}).Where("ID = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count+?", 1))
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
func DeleteComment(ctx context.Context, commentID int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		comment := new(Comment)
		if err := tx.First(&comment, commentID).Error; err != nil {
			return err
		}
		//删除评论，tx.Unscoped().Delete()将永久删除记录
		err := tx.Unscoped().Delete(&comment).Error
		if err != nil {
			return err
		}
		//改变video表中的comment count
		res := tx.Model(&dbFeed.Video{}).Where("ID = ?", vid).Update("comment_count", gorm.Expr("comment_count-?", 1))
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
func GetVideoComments(ctx context.Context, vid int) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{
		VideoID: vid,
	}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
