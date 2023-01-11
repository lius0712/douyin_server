package pack

import (
	"errors"
	"github.com/lius0712/douyin_server/kitex_gen/user"
	"github.com/lius0712/douyin_server/pkg/errno"
)

//封装服务端RPC响应

func BuildUserBaseResp(err error) *user.GetUserResponse {
	if err == nil {
		return userResp(errno.Success)
	}
	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return userResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return userResp(s)

}

func userResp(err errno.ErrNo) *user.GetUserResponse {
	return &user.GetUserResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}

func BuildUserRegisterResp(err error) *user.UseRegisterResponse {
	if err == nil {
		return userRegisterResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return userRegisterResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return userRegisterResp(s)
}

func userRegisterResp(err errno.ErrNo) *user.UseRegisterResponse {
	return &user.UseRegisterResponse{
		StatusCode: err.ErrCode,
		StatusMsg:  &err.ErrMsg,
	}
}

func BuildUserLoginResp(err error) *user.UserLoginResponse {
	if err == nil {
		return userLoginResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return userLoginResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return userLoginResp(s)
}

func userLoginResp(err errno.ErrNo) *user.UserLoginResponse {
	return &user.UserLoginResponse{
		StatusCode: err.ErrCode,
		StatusMsg:  &err.ErrMsg,
	}
}
