package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lius0712/douyin_server/cmd/user/dal"
	user "github.com/lius0712/douyin_server/kitex_gen/user/userservice"
	"github.com/lius0712/douyin_server/pkg/constants"
	"github.com/lius0712/douyin_server/pkg/jwt"
	"net"
)

var Jwt *jwt.JWT

func Init() {
	dal.Init()
	Jwt = jwt.NewJWt([]byte("signingKey"))
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.UserServerAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.UserServerName,
		}),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatalf("userService stopped with error:", err)
	}
}
