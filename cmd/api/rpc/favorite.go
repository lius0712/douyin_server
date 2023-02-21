package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lius0712/douyin_server/kitex_gen/favorite"
	"github.com/lius0712/douyin_server/kitex_gen/favorite/favoriteservice"
	"github.com/lius0712/douyin_server/pkg/constants"
	"github.com/lius0712/douyin_server/pkg/errno"
	"time"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := favoriteservice.NewClient(
		constants.FavoriteServerName,
		client.WithMuxConnection(1),                       //mux
		client.WithRPCTimeout(300*time.Second),            //rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    //conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), //retry
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}
func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp, err = favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp, err = favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
