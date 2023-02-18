package pack

import (
	"context"
	"errors"
	"fmt"
	"github.com/lius0712/douyin_server/cmd/feed/dal/db"
	userDb "github.com/lius0712/douyin_server/cmd/user/dal/db"
	"github.com/lius0712/douyin_server/cmd/user/pack"
	"github.com/lius0712/douyin_server/kitex_gen/feed"
	"gorm.io/gorm"
)

func Video(ctx context.Context, video *db.Video, fromID int64) (*feed.Video, error) {
	if video == nil {
		return nil, nil
	}
	fmt.Println("&&&&&&&&")
	fmt.Println(video.AuthorId)
	authorUser, err := userDb.QueryUserByID(ctx, int64(video.AuthorId))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	author, err := pack.User(ctx, authorUser, fromID)
	if err != nil {
		return nil, err
	}
	return &feed.Video{
		Id:            int64(video.ID),
		Author:        author,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: int64(video.FavoriteCount),
		CommentCount:  int64(video.CommentCount),
		Title:         video.Title,
	}, nil
}
func Videos(ctx context.Context, videos []*db.Video, fromID *int64) ([]*feed.Video, error) {
	authorVideos := make([]*feed.Video, 0)
	for _, video := range videos {
		authorVideo, err := Video(ctx, video, *fromID)
		if err != nil {
			return nil, err
		}
		if authorVideo != nil {
			//TODO:Relation
			authorVideo.IsFavorite = false
			authorVideos = append(authorVideos, authorVideo)
		}
	}
	return authorVideos, nil
}
