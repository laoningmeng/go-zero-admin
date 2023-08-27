package logic

import (
	"context"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/types"
	"github.com/laoningmeng/go-zero-admin/common/encrypt"
	"github.com/laoningmeng/go-zero-admin/services/admin/admin"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginReply, err := l.svcCtx.Rpc.Login(l.ctx, &admin.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	jwtServer := encrypt.NewJwt(encrypt.JwtBaseInfo{
		ExpiresAt: time.Now().Add(3 * time.Hour),
		Secret:    l.svcCtx.Config.JwtAuth.AccessSecret,
	}, types.User{Id: loginReply.UserId})

	token, err := jwtServer.Generate()
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Code:    20000,
		Message: "success",
		Token:   token,
	}, nil
}

func (l *LoginLogic) Logout(token string) (*types.LogoutResp, error) {
	var userInfo types.User
	err := encrypt.GetDataFromToken(token, l.svcCtx.Config.JwtAuth.AccessSecret, &userInfo)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.Rpc.Logout(l.ctx, &admin.LogoutReq{UserId: userInfo.Id})
	if err != nil {
		return nil, err
	}
	return &types.LogoutResp{
		Code:    20000,
		Message: "success",
	}, nil
}
