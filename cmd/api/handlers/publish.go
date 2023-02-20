package handlers

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lius0712/douyin_server/cmd/api/rpc"
	"github.com/lius0712/douyin_server/cmd/publish/pack"
	"github.com/lius0712/douyin_server/kitex_gen/publish"
	"github.com/lius0712/douyin_server/pkg/errno"
	"io"
	"strconv"
)

func Publish(c *gin.Context) {
	var publishParam PublishParam
	token := c.PostForm("token")
	title := c.PostForm("title")
	file, _, err := c.Request.FormFile("data")
	if err != nil {
		SendResponse(c, pack.BuildPublishActionResp(errno.DecodingFailed))
		return
	}
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		SendResponse(c, pack.BuildPublishActionResp(err))
		return
	}
	publishParam.Token = token
	publishParam.Title = title
	resp, err := rpc.PublishAction(context.Background(), &publish.PublishActionRequest{
		File:  buf.Bytes(),
		Token: publishParam.Token,
		Title: publishParam.Title,
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishActionResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}

func PublishList(c *gin.Context) {
	var publishParam UserParam
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ErrHttpBind))
		return
	}
	publishParam.UserId = int64(userId)
	publishParam.Token = c.Query("token")
	if publishParam.UserId < 0 {
		SendResponse(c, pack.BuildPublishListResp(errno.ErrHttpBind))
		return
	}
	resp, err := rpc.PublishList(context.Background(), &publish.PublishListRequest{
		UserId: publishParam.UserId,
		Token:  publishParam.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
