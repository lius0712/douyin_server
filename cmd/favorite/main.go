package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lius0712/douyin_server/cmd/favorite/dal"
	favorite "github.com/lius0712/douyin_server/kitex_gen/favorite/favoriteservice"
	"github.com/lius0712/douyin_server/pkg/constants"
	"github.com/lius0712/douyin_server/pkg/jwt"
	"net"
)

var Jwt *jwt.JWT

func Init() {
	dal.Init()
	Jwt = jwt.NewJWt([]byte(constants.JwtKey))
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.FavoriteServerAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.FavoriteServerName,
		}),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatalf("favoriteService stopped with error:", err)
	}
}
