package model

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/laoningmeng/go-zero-admin/common/logger"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/logic"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id        int32
	Name      string
	Title     string
	Status    int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type RoleModel struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewRoleModel(conn *DB, logger logger.Logger) logic.RoleRepo {
	return &RoleModel{
		db:     conn.db,
		logger: logger,
	}
}

func (u *RoleModel) TableName() string {
	return "role"
}

func (u *RoleModel) Query() *gorm.DB {
	return u.db.Table(u.TableName()).Where(map[string]interface{}{"deleted_at": nil})
}

func (u *RoleModel) FindOne(ctx context.Context, query *logic.Role) (*logic.Role, error) {
	var result Role
	err := u.Query().Where(&Role{
		Id:   query.Id,
		Name: query.Name,
	}).First(&result).Error
	if err != nil {
		return nil, err
	}
	var resp *logic.Role
	_ = copier.Copy(resp, result)
	return resp, nil
}

func (u *RoleModel) Add(ctx context.Context, role *logic.Role) (int32, error) {
	_, err := u.FindOne(ctx, &logic.Role{Name: role.Name})
	switch err {
	case nil:
		return 0, errors.New("用户已经存在不能重复添加")
	case gorm.ErrRecordNotFound:
	default:
		return 0, err
	}
	var data Role
	err = copier.Copy(&data, role)
	if err != nil {
		panic(err)
	}

	err = u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return data.Id, nil
}

func (r *RoleModel) Update(ctx context.Context, filter *logic.Role) (bool, error) {
	role, err := r.FindOne(ctx, filter)
	if err != nil {
		return false, err
	}
	err = r.db.Transaction(func(tx *gorm.DB) error {
		return r.db.Model(&Role{}).Where("id=?", role.Id).Updates(Role{
			Title:  filter.Title,
			Status: filter.Status,
		}).Error
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RoleModel) Delete(ctx context.Context, filter *logic.Role) (bool, error) {
	role, err := r.FindOne(ctx, filter)
	if err != nil {
		return false, err
	}
	err = r.db.Transaction(func(tx *gorm.DB) error {
		return r.db.Model(&Role{}).Where("id=?", role.Id).Update("deleted_at", time.Now()).Error
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *RoleModel) List(ctx context.Context, filter *logic.Role, pageNum, pageSize int) ([]*logic.Role, int64, error) {
	var userList []User
	var total int64
	var queryUser User
	err := copier.Copy(&queryUser, filter)
	if err != nil {
		return nil, 0, err
	}
	u.Query().Where(&queryUser).Preload("Role").Count(&total)

	start := (pageNum - 1) * pageSize

	err = u.Query().Where(&queryUser).Limit(pageSize).Offset(start).Find(&userList).Error

	if err != nil {
		return nil, 0, err
	}
	var result []*logic.Role
	for _, e := range userList {
		var item logic.Role
		err := copier.Copy(&item, e)
		if err != nil {
			return nil, 0, err
		}
		result = append(result, &item)
	}
	return nil, 0, nil
}