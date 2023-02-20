package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lius0712/douyin_server/cmd/api/rpc"
	"github.com/lius0712/douyin_server/cmd/comment/pack"
	"github.com/lius0712/douyin_server/kitex_gen/comment"
	"github.com/lius0712/douyin_server/pkg/errno"
	"strconv"
)

func CommentAction(c *gin.Context) {
	var commentParam CommentActionParam
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	vId, err := strconv.Atoi(videoId)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ConvertErr(err)))
		return
	}
	action, err := strconv.Atoi(actionType)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ConvertErr(err)))
		return
	}
	commentParam.Token = token
	commentParam.VideoId = int64(vId)
	commentParam.ActionType = int32(action)
	rpcRes := comment.CommentActionRequest{
		Token:      commentParam.Token,
		VideoId:    commentParam.VideoId,
		ActionType: commentParam.ActionType,
	}
	if action == 1 {
		commentText := c.Query("comment_text")
		rpcRes.CommentText = &commentText
	} else {
		commentId := c.Query("comment_id")
		cid, err := strconv.Atoi(commentId)
		if err != nil {
			SendResponse(c, pack.BuildCommentActionResp(errno.ConvertErr(err)))
			return
		}
		c := int64(cid)
		rpcRes.CommentId = &c
	}
	resp, err := rpc.CommentAction(context.Background(), &rpcRes)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
func CommentList(c *gin.Context) {
	var commentListParam CommentListParam
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		SendResponse(c, pack.BuildCommentListResp(errno.ConvertErr(err)))
		return
	}
	commentListParam.VideoId = int64(videoId)
	commentListParam.Token = c.Query("token")
	if len(commentListParam.Token) == 0 || commentListParam.VideoId < 0 {
		SendResponse(c, pack.BuildCommentListResp(errno.ErrHttpBind))
		return
	}
	resp, err := rpc.CommentList(context.Background(), &comment.CommentListRequest{
		VideoId: commentListParam.VideoId,
		Token:   commentListParam.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildCommentListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
