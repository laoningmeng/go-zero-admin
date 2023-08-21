package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/types"
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
	token, err := getToken(l.svcCtx.Config.JwtAuth.AccessSecret, loginReply.UserId)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{Token: token}, nil
}

func getToken(secretKey string, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + 60*60*24
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
