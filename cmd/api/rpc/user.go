package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lius0712/douyin_server/kitex_gen/user"
	"github.com/lius0712/douyin_server/kitex_gen/user/userservice"
	"github.com/lius0712/douyin_server/pkg/constants"
	"github.com/lius0712/douyin_server/pkg/errno"
	"time"
)

var userClient userservice.Client

func initUserRpc() {

	//服务发现
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})

	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServerName,
		client.WithMuxConnection(1),                       //mux
		client.WithRPCTimeout(30*time.Second),             //rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), //conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), //retry
		client.WithResolver(r),                            //resolver
	)

	if err != nil {
		panic(err)
	}

	userClient = c
}

func Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UseRegisterResponse, err error) {
	resp, err = userClient.Register(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}

	return resp, nil

}

func Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp, err = userClient.Login(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}

	return resp, nil
}

func GetUserById(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	resp, err = userClient.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, err
}
