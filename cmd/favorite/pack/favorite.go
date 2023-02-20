package pack

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/feed/dal/db"
	"github.com/lius0712/douyin_server/cmd/feed/pack"
	"github.com/lius0712/douyin_server/kitex_gen/feed"
)

func FavoriteVideos(ctx context.Context, vs []db.Video, uid *int64) ([]*feed.Video, error) {
	videos := make([]*db.Video, 0)
	for _, v := range vs {
		videos = append(videos, &v)
	}
	packsVideos, err := pack.Videos(ctx, videos, uid)
	if err != nil {
		return nil, err
	}
	return packsVideos, nil
}
