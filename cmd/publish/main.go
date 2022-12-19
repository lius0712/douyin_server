package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lius0712/douyin_server/cmd/publish/dal"
	publish "github.com/lius0712/douyin_server/kitex_gen/publish/publishservice"
	"github.com/lius0712/douyin_server/pkg/constants"
	"log"
	"net"
)

func Init() {
	dal.Init()
}

func main() {

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.PublishServerAddress)

	if err != nil {
		panic(err)
	}

	Init()

	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.PublishServerName,
		}),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
