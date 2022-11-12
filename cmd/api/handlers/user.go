package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lius0712/douyin_server/cmd/api/rpc"
	"github.com/lius0712/douyin_server/cmd/user/pack"
	"github.com/lius0712/douyin_server/kitex_gen/user"
	"github.com/lius0712/douyin_server/pkg/errno"
	"strconv"
)

func Register(c *gin.Context) {
	var registerParam UserRegisterParam
	registerParam.UserName = c.Query("username")
	registerParam.Password = c.Query("password")

	if len(registerParam.UserName) == 0 || len(registerParam.Password) == 0 {
		SendResponse(c, pack.BuildUserRegisterResp(errno.ErrHttpBind))
		return
	}

	resp, err := rpc.Register(context.Background(), &user.UserRegisterRequest{
		Username: registerParam.UserName,
		Password: registerParam.Password,
	})

	if err != nil {
		SendResponse(c, pack.BuildUserRegisterResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

func Login(c *gin.Context) {
	var loginParam LoginParam
	loginParam.UserName = c.Query("username")
	loginParam.Password = c.Query("password")

	if len(loginParam.UserName) == 0 || len(loginParam.Password) == 0 {
		SendResponse(c, pack.BuildUserLoginResp(errno.ErrHttpBind))
		return
	}

	resp, err := rpc.Login(context.Background(), &user.UserLoginRequest{
		Username: loginParam.UserName,
		Password: loginParam.Password,
	})

	if err != nil {
		SendResponse(c, pack.BuildUserLoginResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)

}

func UserInfo(c *gin.Context) {
	var userParam UserParam
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildUserBaseResp(errno.ErrHttpBind))
		return
	}

	userParam.UserId = int64(userId)
	userParam.Token = c.Query("token")

	//TODO: if Token == 0
	if userParam.UserId < 0 {
		SendResponse(c, pack.BuildUserBaseResp(errno.ErrHttpBind))
		return
	}

	token := user.Token{
		Token:  userParam.Token,
		UserId: userParam.UserId,
	}

	resp, err := rpc.GetUserById(c, &user.GetUserRequest{
		Token: &token,
	})

	if err != nil {
		SendResponse(c, pack.BuildUserBaseResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}
