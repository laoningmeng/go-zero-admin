package logic

import (
	"context"
	"errors"
	"github.com/laoningmeng/go-zero-admin/common/encrypt"
	"github.com/laoningmeng/go-zero-admin/common/logger"
	"github.com/laoningmeng/go-zero-admin/common/trace"
)

// User 这里是逻辑层的User， 跟orm中是有区别的
type User struct {
	Id             int64
	Username       string // 用户名
	Password       string // 密码
	Avatar         string // 头像
	Introduction   string // 个人介绍
	RoleId         int32  // 角色id
	Status         int32  // 0-待激活1-已入职-2-离职中3-已离职
	RoleName       string
	DepartmentName string
	DepartmentId   int32
}

type UserRepo interface {
	FindOne(context.Context, *User) (*User, error)
	TableName() string
	Add(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (bool, error)
	Delete(ctx context.Context, user *User) (bool, error)
	List(ctx context.Context, filter *User, pageNum, pageSize int) ([]*User, int64, error)
}

type UserLogic struct {
	logger logger.Logger
	data   UserRepo
}

func NewUserLogic(repo UserRepo, logger logger.Logger) *UserLogic {
	return &UserLogic{
		data:   repo,
		logger: logger,
	}
}

func (u *UserLogic) Login(ctx context.Context, user *User) (i int64, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-Login]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	queryUser, err := u.data.FindOne(ctx, user)
	if err != nil {
		return 0, err
	}
	if queryUser.Password != encrypt.Md5(user.Password) {
		return 0, errors.New("wrong account or password")
	}
	return queryUser.Id, nil
}
func (u *UserLogic) Logout(ctx context.Context, user *User) (b bool, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-Logout]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	return true, nil
}

func (u *UserLogic) UserQuery(ctx context.Context, user *User) (res *User, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-UserQuery]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	res, err = u.data.FindOne(ctx, user)
	return
}

func (u *UserLogic) UserAdd(ctx context.Context, user *User) (res int64, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-UserAdd]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	res, err = u.data.Add(ctx, user)
	return
}
func (u *UserLogic) UserUpdate(ctx context.Context, user *User) (res bool, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-UserUpdate]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	res, err = u.data.Update(ctx, user)
	return
}

func (u *UserLogic) UserList(ctx context.Context, filter *User, pageNum, pageSize int) (res []*User, total int64, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-UserList]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	return u.data.List(ctx, filter, pageNum, pageSize)
}
func (u *UserLogic) UserDelete(ctx context.Context, user *User) (res bool, err error) {
	ctx, span := trace.StartSpan(ctx, "[user-UserDelete]")
	defer func() {
		trace.EndSpan(span, err)
	}()
	res, err = u.data.Delete(ctx, user)
	return
}
