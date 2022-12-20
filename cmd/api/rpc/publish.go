package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lius0712/douyin_server/kitex_gen/publish"
	"github.com/lius0712/douyin_server/kitex_gen/publish/publishservice"
	"github.com/lius0712/douyin_server/pkg/constants"
	"github.com/lius0712/douyin_server/pkg/errno"
	"time"
)

var publishClient publishservice.Client

func initPublishRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})

	if err != nil {
		panic(err)
	}

	c, err := publishservice.NewClient(
		constants.PublishServerName,
		client.WithMuxConnection(1),                       //mux
		client.WithRPCTimeout(3*time.Second),              //rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    //conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), //retry
		client.WithResolver(r),                            //resolver
	)

	if err != nil {
		panic(err)
	}

	publishClient = c
}

func PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	resp, err = publishClient.PublishAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Status != 0 {
		return nil, errno.NewErrNo(int64(resp.Status), *resp.Msg)
	}
	return resp, nil
}

func PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp, err = publishClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Status != 0 {
		return nil, errno.NewErrNo(int64(resp.Status), *resp.Msg)
	}
	return resp, nil
}
