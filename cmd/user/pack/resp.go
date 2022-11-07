package pack

import (
	"errors"
	"github.com/lius0712/douyin_server/kitex_gen/user"
	"github.com/lius0712/douyin_server/pkg/errno"
)

//封装服务端RPC响应

func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)

}

func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{Status: int32(err.ErrCode), Msg: &err.ErrMsg}
}
