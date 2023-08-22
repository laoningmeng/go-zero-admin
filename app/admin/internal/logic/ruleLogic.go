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

type RuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RuleLogic {
	return &RuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (r *RuleLogic) Add(req *types.RuleAddReq) (*types.RuleAddResp, error) {
	var data admin.RuleAddReq
	err := copier.Copy(&data, req)
	if err != nil {
		panic(err)
	}
	_, err = r.svcCtx.Rpc.RuleAdd(r.ctx, &data)
	if err != nil {
		return nil, err
	}
	return &types.RuleAddResp{
		Code:    0,
		Message: "success",
	}, nil
}

func (u *RuleLogic) List(req *types.RuleListReq) (*types.RuleListResp, error) {
	reply, err := u.svcCtx.Rpc.RuleList(u.ctx, &admin.RuleListReq{
		Name:  req.Name,
		Type:  req.Type,
		Title: req.Title,
		Page: &admin.Page{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	})

	if err != nil {
		return nil, err
	}
	var result types.RuleListResp
	result.Total = reply.Total
	result.Code = 0
	var items []types.Rule
	for _, e := range reply.Data {
		items = append(items, types.Rule{
			Id:        e.Id,
			Name:      e.Name,
			Title:     e.Title,
			Type:      e.Type,
			CreatedAt: time.Unix(e.CreatedAt, 0),
			UpdatedAt: time.Unix(e.UpdatedAt, 0),
		})
	}
	result.Items = items

	return &result, nil
}

func (u *RuleLogic) Update(req *types.RuleUpdateReq) (*types.RuleUpdateResp, error) {
	_, err := u.svcCtx.Rpc.RuleUpdate(u.ctx, &admin.RuleUpdateReq{
		Id:    req.Id,
		Name:  req.Name,
		Title: req.Title,
		Type:  req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &types.RuleUpdateResp{
		Code:    0,
		Message: "success",
	}, nil
}

func (u *RuleLogic) Del(req *types.RuleDelReq) (*types.RuleDelResp, error) {
	_, err := u.svcCtx.Rpc.RuleDelete(u.ctx, &admin.RuleDeleteReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.RuleDelResp{
		Code:    0,
		Message: "success",
	}, nil
}
