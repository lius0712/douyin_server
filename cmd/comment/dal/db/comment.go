package db

import (
	dbFeed "github.com/lius0712/douyin_server/cmd/feed/dal/db"
	dbUser "github.com/lius0712/douyin_server/cmd/user/dal/db"
	"github.com/lius0712/douyin_server/pkg/constants"
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
