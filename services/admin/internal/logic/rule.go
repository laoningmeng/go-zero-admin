package logic

import (
	"context"
	"github.com/laoningmeng/go-zero-admin/common/logger"
	"github.com/laoningmeng/go-zero-admin/common/trace"
	"time"
)

// User 这里是逻辑层的User， 跟orm中是有区别的
type Rule struct {
	Id        int32
	Name      string // 用户名
	Title     string // title
	Type      int32  // 1-menu 2-btn
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

func (r *RuleLogic) RuleQuery(ctx context.Context, role *Rule) (res *Rule, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-RoleQuery]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	res, err = r.data.FindOne(ctx, role)
	return
}

func (r *RuleLogic) RuleAdd(ctx context.Context, role *Rule) (res int32, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-RuleAdd]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	res, err = r.data.Add(ctx, role)
	return
}
func (r *RuleLogic) RuleUpdate(ctx context.Context, role *Rule) (res bool, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-RuleUpdate]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	res, err = r.data.Update(ctx, role)
	return
}

func (r *RuleLogic) RuleList(ctx context.Context, filter *Rule, pageNum, pageSize int) (res []*Rule, total int32, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-RuleList]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	res, total, err = r.data.List(ctx, filter, pageNum, pageSize)
	return
}
func (r *RuleLogic) RuleDelete(ctx context.Context, user *Rule) (res bool, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-RuleDelete]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	res, err = r.data.Delete(ctx, user)
	return
}
