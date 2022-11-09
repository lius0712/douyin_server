package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lius0712/douyin_server/cmd/api/rpc"
	"github.com/lius0712/douyin_server/cmd/user/pack"
	"github.com/lius0712/douyin_server/kitex_gen/user"
	"github.com/lius0712/douyin_server/pkg/errno"
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
