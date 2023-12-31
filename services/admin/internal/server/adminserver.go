// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package server

import (
	"github.com/google/wire"

	"github.com/laoningmeng/go-zero-admin/services/admin/admin"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/logic"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/svc"
)

var ProviderSet = wire.NewSet(NewAdminServer)

type AdminServer struct {
	svcCtx *svc.ServiceContext
	u      *logic.UserLogic // user
	r      *logic.RoleLogic // role
	a      *logic.RuleLogic // auth
	admin.UnimplementedAdminServer
}

func NewAdminServer(svcCtx *svc.ServiceContext, userLogic *logic.UserLogic, roleLogic *logic.RoleLogic, ruleLogic *logic.RuleLogic) *AdminServer {
	return &AdminServer{
		svcCtx: svcCtx,
		u:      userLogic,
		r:      roleLogic,
		a:      ruleLogic,
	}
}
