package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/types"
	"github.com/laoningmeng/go-zero-admin/services/admin/admin"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type RoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleLogic {
	return &RoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (r *RoleLogic) Add(req *types.RoleAddReq) (*types.RoleAddResp, error) {
	var data admin.RoleAddReq
	err := copier.Copy(&data, req)
	if err != nil {
		panic(err)
	}
	_, err = r.svcCtx.Rpc.RoleAdd(r.ctx, &data)
	if err != nil {
		return nil, err
	}
	return &types.RoleAddResp{
		Code:    0,
		Message: "success",
	}, nil
}

func (u *RoleLogic) List(req *types.RoleListReq) (*types.RoleListResp, error) {
	reply, err := u.svcCtx.Rpc.RoleList(u.ctx, &admin.RoleListReq{
		Name:   req.Name,
		Status: req.Status,
		Title:  req.Title,
		Page: &admin.Page{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	})

	if err != nil {
		return nil, err
	}
	var result types.RoleListResp
	result.Total = reply.Total
	result.Code = 0
	var items []types.Role
	for _, e := range reply.Data {
		items = append(items, types.Role{
			Id:        e.Id,
			Name:      e.Name,
			Title:     e.Title,
			Status:    e.Status,
			CreatedAt: time.Unix(e.CreatedAt, 0),
			UpdatedAt: time.Unix(e.UpdatedAt, 0),
		})
	}
	result.Items = items

	return &result, nil
}

func (u *RoleLogic) Update(req *types.RoleUpdateReq) (*types.RoleUpdateResp, error) {
	_, err := u.svcCtx.Rpc.RoleUpdate(u.ctx, &admin.RoleUpdateReq{
		Id:     req.Id,
		Name:   req.Name,
		Title:  req.Title,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.RoleUpdateResp{
		Code:    0,
		Message: "success",
	}, nil
}

func (u *RoleLogic) Del(req *types.RoleDelReq) (*types.RoleDelResp, error) {
	_, err := u.svcCtx.Rpc.RoleDelete(u.ctx, &admin.RoleDeleteReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.RoleDelResp{
		Code:    0,
		Message: "success",
	}, nil
}
