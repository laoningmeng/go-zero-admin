package model

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/laoningmeng/go-zero-admin/common/encrypt"
	"github.com/laoningmeng/go-zero-admin/common/logger"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/logic"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id           int64
	Username     string
	Password     string
	Avatar       string
	Introduction string
	RoleId       int32
	DepartmentId int32
	Status       int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Role         Role
	Department   Department
}

type UserModel struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewUserModel(conn *DB, logger logger.Logger) logic.UserRepo {
	return &UserModel{
		db:     conn.db,
		logger: logger,
	}
}

func (u *UserModel) TableName() string {
	return "user"
}

func (u *UserModel) Query() *gorm.DB {
	return u.db.Table(u.TableName()).Where(map[string]interface{}{"deleted_at": nil})
}

func (u *UserModel) FindOne(ctx context.Context, query *logic.User) (*logic.User, error) {
	var Result User
	err := u.Query().Where(&User{
		Id:       query.Id,
		Username: query.Username,
	}).Preload("Role").Preload("Department").First(&Result).Error
	if err != nil {
		return nil, err
	}
	return &logic.User{
		Id:             Result.Id,
		Username:       Result.Username,
		Password:       Result.Password,
		Avatar:         Result.Avatar,
		Introduction:   Result.Introduction,
		RoleId:         Result.RoleId,
		Status:         Result.Status,
		RoleName:       Result.Role.Title,
		DepartmentId:   Result.DepartmentId,
		DepartmentName: Result.Department.Name,
	}, nil
}

func (u *UserModel) Add(ctx context.Context, user *logic.User) (int64, error) {
	_, err := u.FindOne(ctx, &logic.User{Username: user.Username})
	switch err {
	case nil:
		return 0, errors.New("用户已经存在不能重复添加")
	case gorm.ErrRecordNotFound:
	default:
		return 0, err
	}
	data := User{
		Username:     user.Username,
		Password:     encrypt.Md5("888888"),
		Avatar:       user.Avatar,
		Introduction: user.Introduction,
		RoleId:       user.RoleId,
		Status:       user.Status,
		DepartmentId: user.DepartmentId,
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

func (u *UserModel) Update(ctx context.Context, filter *logic.User) (bool, error) {
	user, err := u.FindOne(ctx, &logic.User{
		Id:       filter.Id,
		Username: filter.Username,
	})
	if err != nil {
		return false, err
	}
	var update User
	err = copier.Copy(&update, filter)
	if err != nil {
		return false, err
	}
	err = u.db.Transaction(func(tx *gorm.DB) error {
		return u.db.Model(&User{}).Where("id=?", user.Id).Updates(&update).Error
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserModel) Delete(ctx context.Context, filter *logic.User) (bool, error) {
	user, err := u.FindOne(ctx, filter)
	if err != nil {
		return false, err
	}
	err = u.db.Transaction(func(tx *gorm.DB) error {
		return u.db.Model(&User{}).Where("id=?", user.Id).Update("deleted_at", time.Now()).Error
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserModel) List(ctx context.Context, filter *logic.User, pageNum, pageSize int) ([]*logic.User, int64, error) {
	var userList []User
	var total int64
	u.Query().Where(&User{
		Username: filter.Username,
		RoleId:   filter.RoleId,
		Status:   filter.Status,
	}).Preload("Role").Count(&total)

	start := (pageNum - 1) * pageSize

	err := u.Query().Where(&User{
		Username:     filter.Username,
		RoleId:       filter.RoleId,
		Status:       filter.Status,
		DepartmentId: filter.DepartmentId,
	}).Preload("Role").Preload("Department").Limit(pageSize).Offset(start).Find(&userList).Error

	if err != nil {
		return nil, 0, err
	}
	var result []*logic.User
	for _, e := range userList {
		result = append(result, &logic.User{
			Id:             e.Id,
			Username:       e.Username,
			Avatar:         e.Avatar,
			Introduction:   e.Introduction,
			RoleId:         e.RoleId,
			Status:         e.Status,
			RoleName:       e.Role.Title,
			DepartmentId:   e.DepartmentId,
			DepartmentName: e.Department.Name,
		})
	}
	return result, total, nil
}
