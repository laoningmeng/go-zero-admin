package logic

import (
	"context"
	"github.com/laoningmeng/go-zero-admin/common/logger"
	"time"
)

// User 这里是逻辑层的User， 跟orm中是有区别的
type Role struct {
	Id        int32
	Name      string // 用户名
	Title     string // title
	Status    int32  // 0-待激活1-已入职-2-离职中3-已离职
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type RoleRepo interface {
	FindOne(context.Context, *Role) (*Role, error)
	TableName() string
	Add(ctx context.Context, user *Role) (int32, error)
	Update(ctx context.Context, user *Role) (bool, error)
	Delete(ctx context.Context, user *Role) (bool, error)
	List(ctx context.Context, filter *Role, pageNum, pageSize int) ([]*Role, int32, error)
}

type RoleLogic struct {
	logger logger.Logger
	data   RoleRepo
}

func NewRoleLogic(repo RoleRepo, logger logger.Logger) *RoleLogic {
	return &RoleLogic{
		data:   repo,
		logger: logger,
	}
}

func (r *RoleLogic) RoleQuery(ctx context.Context, role *Role) (*Role, error) {
	return r.data.FindOne(ctx, role)
}

func (r *RoleLogic) RoleAdd(ctx context.Context, role *Role) (int32, error) {
	return r.data.Add(ctx, role)
}
func (r *RoleLogic) RoleUpdate(ctx context.Context, role *Role) (bool, error) {
	return r.data.Update(ctx, role)
}

func (r *RoleLogic) RoleList(ctx context.Context, filter *Role, pageNum, pageSize int) ([]*Role, int32, error) {
	return r.data.List(ctx, filter, pageNum, pageSize)
}
func (r *RoleLogic) RoleDelete(ctx context.Context, user *Role) (bool, error) {
	return r.data.Delete(ctx, user)
}
