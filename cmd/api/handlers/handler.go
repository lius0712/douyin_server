package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

type UserRegisterParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// 用户信息输出参数
type UserParam struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type PublishParam struct {
	Data  []byte `json:"data,omitempty"`
	Token string `json:"token,omitempty"`
	Title string `json:"title,omitempty"`
}

type FeedParam struct {
	LatestTime *int64  `json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `json:"token,omitempty"`       // 可选参数，登录用户设置
}
