package main

import (
	"context"
	"fmt"
	"github.com/lius0712/douyin_server/cmd/user/pack"
	"github.com/lius0712/douyin_server/cmd/user/service"
	user "github.com/lius0712/douyin_server/kitex_gen/user"
	"github.com/lius0712/douyin_server/pkg/errno"
	"github.com/lius0712/douyin_server/pkg/jwt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UseRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UseRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserRegisterResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateService(ctx).CreateUser(req)

	if err != nil {
		resp = pack.BuildUserRegisterResp(err)
	}

	resp = pack.BuildUserRegisterResp(errno.Success)

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...

	resp = new(user.UserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserLoginResp(errno.ErrHttpBind)
		return resp, nil
	}
	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuildUserLoginResp(err)
		return resp, err
	}

	token, err := Jwt.CreateToken(jwt.MyCustomClaims{
		Id: uid,
	})

	if err != nil {
		resp = pack.BuildUserLoginResp(errno.ErrorTokenInvalid)
		return resp, err
	}

	resp = pack.BuildUserLoginResp(errno.Success)
	resp.Token = &user.Token{UserId: uid, Token: token}
	fmt.Println("*******")
	fmt.Println(resp)
	return resp, nil
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	// TODO: Your code here...

	return
}
