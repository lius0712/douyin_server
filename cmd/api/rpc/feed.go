package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lius0712/douyin_server/kitex_gen/feed"
	"github.com/lius0712/douyin_server/kitex_gen/feed/feedservice"
	"github.com/lius0712/douyin_server/pkg/constants"
	"github.com/lius0712/douyin_server/pkg/errno"
	"time"
)

var feedClient feedservice.Client

func initFeedRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := feedservice.NewClient(
		constants.FeedServerName,
		client.WithMuxConnection(1),                       //mux
		client.WithRPCTimeout(300*time.Second),            //rpc timeout
		client.WithConnectTimeout(300*time.Millisecond),   //conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), //retry
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	feedClient = c
}
func GetUserFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	resp, err = feedClient.GetUserFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
