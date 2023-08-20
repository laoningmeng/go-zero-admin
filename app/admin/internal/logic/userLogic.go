package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/types"
	"github.com/laoningmeng/go-zero-admin/services/admin/admin"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *UserLogic) Add(req *types.UserAddReq) (*types.UserAddResp, error) {
	var data admin.UserAddReq
	err := copier.Copy(&data, req)
	if err != nil {
		panic(err)
	}
	_, err = u.svcCtx.UserRpc.UserAdd(u.ctx, &data)
	if err != nil {
		return nil, err
	}
	return &types.UserAddResp{
		Code:    0,
		Message: "success",
	}, nil
}

func (u *UserLogic) List(req *types.UserListReq) (*types.UserListResp, error) {
	reply, err := u.svcCtx.UserRpc.UserList(u.ctx, &admin.UserListReq{
		Username:     req.Username,
		RoleId:       req.RoleId,
		Status:       req.Status,
		DepartmentId: req.DepartmentId,
		Page: &admin.Page{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	})

	if err != nil {
		return nil, err
	}
	var result types.UserListResp
	result.Total = reply.Total
	result.Code = 0
	var items []types.User

	for _, e := range reply.Data {
		fmt.Println("E:", e.Department, e.DepartmentId)
		items = append(items, types.User{
			Id:           e.Id,
			Username:     e.Username,
			RoleName:     e.RoleName,
			RoleId:       e.RoleId,
			DepartmentId: e.DepartmentId,
			Department:   e.Department,
		})
	}
	result.Items = items

	return &result, nil
}

func (u *UserLogic) Update(req *types.UserUpdateReq) (*types.UserUpdateResp, error) {
	_, err := u.svcCtx.UserRpc.UserUpdate(u.ctx, &admin.UserUpdateReq{
		Id:           req.Id,
		RoleId:       req.RoleId,
		DepartmentId: req.DepartmentId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserUpdateResp{
		Code:    0,
		Message: "success",
	}, nil
}

func (u *UserLogic) Del(req *types.UserDelReq) (*types.UserDelResp, error) {
	_, err := u.svcCtx.UserRpc.UserDelete(u.ctx, &admin.UserDeleteReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserDelResp{
		Code:    0,
		Message: "success",
	}, nil
}
