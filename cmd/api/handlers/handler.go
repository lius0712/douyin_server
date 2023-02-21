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

type CommentActionParam struct {
	UserId      int64   `json:"user_id,omitempty"`
	Token       string  `json:"token,omitempty"`
	VideoId     int64   `json:"video_id,omitempty"`
	ActionType  int32   `json:"action_type,omitempty"`  // 1-发布评论，2-删除评论
	CommentText *string `json:"comment_text,omitempty"` //用户填写的评论内容，在action_type=1的时候使用
	CommentId   *int64  `json:"comment_id,omitempty"`   //要删除的评论id，在action_type=2的时候使用
}

type CommentListParam struct {
	Token   string `json:"token,omitempty"`
	VideoId int64  `json:"video_id,omitempty"`
}

type FavoriteActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
	VideoId    int64  `json:"video_id,omitempty"`
	ActionType int32  `json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

type FavoriteListParam struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}
