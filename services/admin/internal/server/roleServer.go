package server

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/laoningmeng/go-zero-admin/services/admin/admin"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/logic"
)

func (s *AdminServer) RoleAdd(ctx context.Context, req *admin.RoleAddReq) (*admin.RoleAddReply, error) {
	var param *logic.Role
	_ = copier.Copy(param, req)
	userId, err := s.r.RoleAdd(ctx, param)
	if err != nil {
		return nil, err
	}
	return &admin.RoleAddReply{
		Id: userId,
	}, nil
}

func (s *AdminServer) RoleUpdate(ctx context.Context, req *admin.RoleUpdateReq) (*admin.RoleUpdateReply, error) {
	var param *logic.Role
	_ = copier.Copy(param, req)

	role, err := s.r.RoleUpdate(ctx, param)

	if err != nil {
		return nil, err
	}
	var resp *admin.RoleUpdateReply
	_ = copier.Copy(resp, role)

	return resp, nil
}

func (s *AdminServer) RoleQuery(ctx context.Context, req *admin.UserQueryReq) (*admin.UserQueryReply, error) {
	user, err := s.u.UserQuery(ctx, &logic.User{
		Id:       0,
		Username: "",
	})
	if err != nil {
		return nil, err
	}
	return &admin.UserQueryReply{
		Id:             user.Id,
		Username:       user.Username,
		Avatar:         user.Avatar,
		Introduction:   user.Introduction,
		RoleId:         user.RoleId,
		RoleName:       user.RoleName,
		Status:         user.Status,
		DepartmentName: user.DepartmentName,
		DepartmentId:   user.DepartmentId,
	}, nil
}

func (s *AdminServer) RoleList(ctx context.Context, req *admin.UserListReq) (*admin.UserListReply, error) {
	list, total, err := s.u.UserList(ctx, &logic.User{
		Username: "",
		RoleId:   0,
		Status:   0,
	}, int(req.Page.Page), int(req.Page.PageSize))
	if err != nil {
		return nil, err
	}
	var data []*admin.UserListReplyDetail
	for _, e := range list {
		data = append(data, &admin.UserListReplyDetail{
			Id:           e.Id,
			Username:     e.Username,
			Avatar:       e.Avatar,
			Introduction: e.Introduction,
			RoleId:       e.RoleId,
			RoleName:     e.RoleName,
			Status:       e.Status,
			Department:   e.DepartmentName,
			DepartmentId: e.DepartmentId,
		})
	}
	return &admin.UserListReply{Total: total, Data: data}, nil
}

func (s *AdminServer) RoleDelete(ctx context.Context, req *admin.UserDeleteReq) (*admin.UserDeleteReply, error) {
	isDel, err := s.u.UserDelete(ctx, &logic.User{
		Id:       req.Id,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	return &admin.UserDeleteReply{
		IsOk: isDel,
	}, nil
}
