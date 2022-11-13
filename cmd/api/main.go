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

	//apiRouter.GET("/feed/", handlers.Feed)
	apiRouter.GET("/user/", handlers.UserInfo)
	apiRouter.POST("/user/register/", handlers.Register)
	apiRouter.POST("/user/login/", handlers.Login)
	apiRouter.POST("/publish/action/", handlers.Publish)
	//apiRouter.GET("/publish/list/", handlers.PublishList)

	// extra apis - I
	//apiRouter.POST("/favorite/action/", handlers.FavoriteAction)
	//apiRouter.GET("/favorite/list/", handlers.FavoriteList)
	//apiRouter.POST("/comment/action/", handlers.CommentAction)
	//apiRouter.GET("/comment/list/", handlers.CommentList)

	// extra apis - II
	//apiRouter.POST("/relation/action/", handlers.RelationAction)
	//apiRouter.GET("/relation/follow/list/", handlers.FollowList)
	//apiRouter.GET("/relation/follower/list/", handlers.FollowerList)

	if err := http.ListenAndServe("127.0.0.1:8088", r); err != nil {
		klog.Fatal(err)
	}
}
