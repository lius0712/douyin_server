package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"github.com/lius0712/douyin_server/cmd/api/handlers"
	"github.com/lius0712/douyin_server/cmd/api/rpc"
	"net/http"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()

	r := gin.New()
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	//apiRouter.GET("/feed/", controller.Feed)
	//apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", handlers.Register)
	//apiRouter.POST("/user/login/", controller.Login)
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//apiRouter.POST("/comment/action/", controller.CommentAction)
	//apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)

	if err := http.ListenAndServe("127.0.0.1:8088", r); err != nil {
		klog.Fatal(err)
	}
}
