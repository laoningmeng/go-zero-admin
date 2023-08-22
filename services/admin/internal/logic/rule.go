package logic

import (
	"context"
	"github.com/laoningmeng/go-zero-admin/common/logger"
	"time"
)

// User 这里是逻辑层的User， 跟orm中是有区别的
type Rule struct {
	Id        int32
	Name      string // 用户名
	Title     string // title
	Type      int32  // 0-待激活1-已入职-2-离职中3-已离职
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type RuleRepo interface {
	FindOne(context.Context, *Rule) (*Rule, error)
	TableName() string
	Add(ctx context.Context, user *Rule) (int32, error)
	Update(ctx context.Context, user *Rule) (bool, error)
	Delete(ctx context.Context, user *Rule) (bool, error)
	List(ctx context.Context, filter *Rule, pageNum, pageSize int) ([]*Rule, int32, error)
}

type RuleLogic struct {
	logger logger.Logger
	data   RuleRepo
}

func NewRuleLogic(repo RuleRepo, logger logger.Logger) *RuleLogic {
	return &RuleLogic{
		data:   repo,
		logger: logger,
	}
}

func (r *RuleLogic) RuleQuery(ctx context.Context, role *Rule) (*Rule, error) {
	return r.data.FindOne(ctx, role)
}

func (r *RuleLogic) RuleAdd(ctx context.Context, role *Rule) (int32, error) {
	return r.data.Add(ctx, role)
}
func (r *RuleLogic) RuleUpdate(ctx context.Context, role *Rule) (bool, error) {
	return r.data.Update(ctx, role)
}

func (r *RuleLogic) RuleList(ctx context.Context, filter *Rule, pageNum, pageSize int) ([]*Rule, int32, error) {
	return r.data.List(ctx, filter, pageNum, pageSize)
}
func (r *RuleLogic) RuleDelete(ctx context.Context, user *Rule) (bool, error) {
	return r.data.Delete(ctx, user)
}
