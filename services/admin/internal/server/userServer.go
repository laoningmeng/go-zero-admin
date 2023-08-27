package server

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/laoningmeng/go-zero-admin/services/admin/admin"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/logic"
)

func (s *AdminServer) Login(ctx context.Context, req *admin.LoginReq) (*admin.LoginReply, error) {
	userId, err := s.u.Login(ctx, &logic.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &admin.LoginReply{UserId: userId}, nil
}

func (s *AdminServer) Logout(ctx context.Context, req *admin.LogoutReq) (*admin.LogoutReply, error) {
	isOk, err := s.u.Logout(ctx, &logic.User{
		Id: req.UserId,
	})
	if err != nil {
		return &admin.LogoutReply{
			IsOk: false,
		}, err
	}
	return &admin.LogoutReply{
		IsOk: isOk,
	}, nil
}

func (s *AdminServer) UserAdd(ctx context.Context, req *admin.UserAddReq) (*admin.UserAddReply, error) {
	userId, err := s.u.UserAdd(ctx, &logic.User{
		Username:     req.Username,
		Password:     req.Password,
		Avatar:       req.Avatar,
		Introduction: req.Introduction,
		RoleId:       req.RoleId,
		DepartmentId: req.DepartmentId,
	})
	if err != nil {
		return nil, err
	}
	return &admin.UserAddReply{
		Id: userId,
	}, nil
}

func (s *AdminServer) UserUpdate(ctx context.Context, req *admin.UserUpdateReq) (*admin.UserUpdateReply, error) {
	_, err := s.u.UserUpdate(ctx, &logic.User{
		Id:           req.Id,
		Username:     req.Username,
		Password:     req.Password,
		Avatar:       req.Avatar,
		Introduction: req.Introduction,
		RoleId:       req.RoleId,
		Status:       req.Status,
	})
	if err != nil {
		return &admin.UserUpdateReply{IsOk: false}, err
	}
	return &admin.UserUpdateReply{
		IsOk: true,
	}, nil
}

func (s *AdminServer) UserQuery(ctx context.Context, req *admin.UserQueryReq) (*admin.UserQueryReply, error) {
	var lu logic.User
	_ = copier.Copy(&lu, req)
	user, err := s.u.UserQuery(ctx, &lu)
	if err != nil {
		return nil, err
	}
	var reply admin.UserQueryReply
	_ = copier.Copy(&reply, user)
	return &reply, nil
}

func (s *AdminServer) UserList(ctx context.Context, req *admin.UserListReq) (*admin.UserListReply, error) {
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

func (s *AdminServer) UserDelete(ctx context.Context, req *admin.UserDeleteReq) (*admin.UserDeleteReply, error) {
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
