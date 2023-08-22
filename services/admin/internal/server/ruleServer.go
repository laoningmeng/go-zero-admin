package server

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/laoningmeng/go-zero-admin/services/admin/admin"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/logic"
)

func (s *AdminServer) RuleAdd(ctx context.Context, req *admin.RuleAddReq) (*admin.RuleAddReply, error) {
	var param logic.Rule
	_ = copier.Copy(&param, req)
	userId, err := s.a.RuleAdd(ctx, &param)
	if err != nil {
		return nil, err
	}
	return &admin.RuleAddReply{
		Id: userId,
	}, nil
}

func (s *AdminServer) RuleUpdate(ctx context.Context, req *admin.RuleUpdateReq) (*admin.RuleUpdateReply, error) {
	var param logic.Rule
	_ = copier.Copy(&param, req)

	role, err := s.a.RuleUpdate(ctx, &param)

	if err != nil {
		return nil, err
	}
	var resp admin.RuleUpdateReply
	_ = copier.Copy(&resp, role)

	return &resp, nil
}

func (s *AdminServer) RuleQuery(ctx context.Context, req *admin.RuleQueryReq) (*admin.RuleQueryReply, error) {
	var fileter logic.Rule
	err := copier.Copy(&fileter, req)
	if err != nil {
		return nil, err
	}
	user, err := s.a.RuleQuery(ctx, &fileter)
	if err != nil {
		return nil, err
	}
	var result admin.RuleQueryReply
	err = copier.Copy(&result, user)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AdminServer) RuleList(ctx context.Context, req *admin.RuleListReq) (*admin.RuleListReply, error) {
	var filter logic.Rule
	err := copier.Copy(&filter, req)
	if err != nil {
		return nil, err
	}
	list, total, err := s.a.RuleList(ctx, &filter, int(req.Page.Page), int(req.Page.PageSize))
	if err != nil {
		return nil, err
	}
	var data []*admin.RuleListReplyItem
	for _, e := range list {
		data = append(data, &admin.RuleListReplyItem{
			Id:        e.Id,
			Name:      e.Name,
			Title:     e.Title,
			Type:      e.Type,
			CreatedAt: e.CreatedAt.Unix(),
			UpdatedAt: e.UpdatedAt.Unix(),
		})
	}
	return &admin.RuleListReply{
		Total: total,
		Data:  data,
	}, nil
}

func (s *AdminServer) RuleDelete(ctx context.Context, req *admin.RuleDeleteReq) (*admin.RuleDeleteReply, error) {
	isDel, err := s.a.RuleDelete(ctx, &logic.Rule{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &admin.RuleDeleteReply{
		IsOk: isDel,
	}, nil
}
