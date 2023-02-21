package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lius0712/douyin_server/cmd/api/rpc"
	"github.com/lius0712/douyin_server/cmd/favorite/pack"
	"github.com/lius0712/douyin_server/kitex_gen/favorite"
	"github.com/lius0712/douyin_server/pkg/errno"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	var favoriteActParam FavoriteActionParam
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	vid, err := strconv.Atoi(videoId)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(actionType)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}
	favoriteActParam.Token = token
	favoriteActParam.VideoId = int64(vid)
	favoriteActParam.ActionType = int32(act)
	resp, err := rpc.FavoriteAction(context.Background(), &favorite.FavoriteActionRequest{
		VideoId:    favoriteActParam.VideoId,
		Token:      favoriteActParam.Token,
		ActionType: favoriteActParam.ActionType,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
func FavoriteList(c *gin.Context) {
	var favoriteListParam FavoriteListParam
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	favoriteListParam.UserId = int64(userId)
	favoriteListParam.Token = c.Query("token")
	if len(favoriteListParam.Token) == 0 || favoriteListParam.UserId < 0 {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	resp, err := rpc.FavoriteList(context.Background(), &favorite.FavoriteListRequest{
		UserId: favoriteListParam.UserId,
		Token:  favoriteListParam.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
