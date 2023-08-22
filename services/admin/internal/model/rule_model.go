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

type Rule struct {
	Id        int32
	Name      string
	Title     string
	Type      int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type RuleModel struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewRuleModel(conn *DB, logger logger.Logger) logic.RuleRepo {
	return &RuleModel{
		db:     conn.db,
		logger: logger,
	}
}

func (u *RuleModel) TableName() string {
	return "rule"
}

func (u *RuleModel) Query() *gorm.DB {
	return u.db.Table(u.TableName()).Where(map[string]interface{}{"deleted_at": nil})
}

func (u *RuleModel) FindOne(ctx context.Context, query *logic.Rule) (*logic.Rule, error) {
	var result Rule
	err := u.Query().Where(&Rule{
		Id:   query.Id,
		Name: query.Name,
	}).First(&result).Error
	if err != nil {
		return nil, err
	}
	var resp logic.Rule
	_ = copier.Copy(&resp, result)
	return &resp, nil
}

func (u *RuleModel) Add(ctx context.Context, role *logic.Rule) (int32, error) {
	_, err := u.FindOne(ctx, &logic.Rule{Name: role.Name})
	switch err {
	case nil:
		return 0, errors.New("用户已经存在不能重复添加")
	case gorm.ErrRecordNotFound:
	default:
		return 0, err
	}
	var data Rule
	err = copier.Copy(&data, role)
	if err != nil {
		panic(err)
		return 0, err
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

func (r *RuleModel) Update(ctx context.Context, filter *logic.Rule) (bool, error) {
	_, err := r.FindOne(ctx, &logic.Rule{Id: filter.Id})
	if err != nil {
		return false, err
	}
	err = r.db.Transaction(func(tx *gorm.DB) error {
		return r.db.Model(&Rule{}).Where("id=?", filter.Id).Updates(Rule{
			Title: filter.Title,
			Type:  filter.Type,
			Name:  filter.Name,
		}).Error
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RuleModel) Delete(ctx context.Context, filter *logic.Rule) (bool, error) {
	role, err := r.FindOne(ctx, filter)
	if err != nil {
		return false, err
	}
	err = r.db.Transaction(func(tx *gorm.DB) error {
		return r.db.Model(&Rule{}).Where("id=?", role.Id).Update("deleted_at", time.Now()).Error
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RuleModel) List(ctx context.Context, filter *logic.Rule, pageNum, pageSize int) ([]*logic.Rule, int32, error) {
	var roleList []Rule
	var total int64
	var queryRule Rule
	if filter != nil {
		err := copier.Copy(&queryRule, filter)
		if err != nil {
			return nil, 0, err
		}
	}

	r.Query().Where(&queryRule).Count(&total)
	start := (pageNum - 1) * pageSize
	err := r.Query().Where(&queryRule).Limit(pageSize).Offset(start).Find(&roleList).Error

	if err != nil {
		return nil, 0, err
	}
	var result []*logic.Rule
	for _, e := range roleList {
		var item logic.Rule
		err := copier.Copy(&item, e)
		if err != nil {
			return nil, 0, err
		}
		result = append(result, &item)
	}
	return result, int32(total), nil
}
