package feed

import (
	"context"
	feed "github.com/lius0712/douyin_server/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetUserFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetUserFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	return
}
