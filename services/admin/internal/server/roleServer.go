package server

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/laoningmeng/go-zero-admin/services/admin/admin"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/logic"
)

func (s *AdminServer) RoleAdd(ctx context.Context, req *admin.RoleAddReq) (*admin.RoleAddReply, error) {
	var param logic.Role
	_ = copier.Copy(&param, req)
	userId, err := s.r.RoleAdd(ctx, &param)
	if err != nil {
		return nil, err
	}
	return &admin.RoleAddReply{
		Id: userId,
	}, nil
}

func (s *AdminServer) RoleUpdate(ctx context.Context, req *admin.RoleUpdateReq) (*admin.RoleUpdateReply, error) {
	var param logic.Role
	_ = copier.Copy(&param, req)

	role, err := s.r.RoleUpdate(ctx, &param)

	if err != nil {
		return nil, err
	}
	var resp admin.RoleUpdateReply
	_ = copier.Copy(&resp, role)

	return &resp, nil
}

func (s *AdminServer) RoleQuery(ctx context.Context, req *admin.RoleQueryReq) (*admin.RoleQueryReply, error) {
	var fileter logic.Role
	err := copier.Copy(&fileter, req)
	if err != nil {
		return nil, err
	}
	user, err := s.r.RoleQuery(ctx, &fileter)
	if err != nil {
		return nil, err
	}
	var result admin.RoleQueryReply
	err = copier.Copy(&result, user)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AdminServer) RoleList(ctx context.Context, req *admin.RoleListReq) (*admin.RoleListReply, error) {
	var filter logic.Role
	err := copier.Copy(&filter, req)
	if err != nil {
		return nil, err
	}
	list, total, err := s.r.RoleList(ctx, &filter, int(req.Page.Page), int(req.Page.PageSize))
	if err != nil {
		return nil, err
	}
	var data []*admin.RoleListReplyItem
	for _, e := range list {
		data = append(data, &admin.RoleListReplyItem{
			Id:        e.Id,
			Name:      e.Name,
			Title:     e.Title,
			Status:    e.Status,
			CreatedAt: e.CreatedAt.Unix(),
			UpdatedAt: e.UpdatedAt.Unix(),
		})
	}
	return &admin.RoleListReply{
		Total: total,
		Data:  data,
	}, nil
}

func (s *AdminServer) RoleDelete(ctx context.Context, req *admin.RoleDeleteReq) (*admin.RoleDeleteReply, error) {
	isDel, err := s.r.RoleDelete(ctx, &logic.Role{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &admin.RoleDeleteReply{
		IsOk: isDel,
	}, nil
}
