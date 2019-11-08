package sys

import (
	"time"

	"iris-ticket/backend/models/basemodel"
	"iris-ticket/backend/models/db"

	"github.com/jinzhu/gorm"
)

// 角色
type Role struct {
	basemodel.Model
	Name     string `gorm:"column:name;size:32;not null;" json:"name" form:"name"`        // 名称
	Sequence int    `gorm:"column:sequence;not null;" json:"sequence" form:"sequence"`    // 排序值
	ParentID uint64 `gorm:"column:parent_id;not null;" json:"parent_id" form:"parent_id"` // 父级ID
	Memo     string `gorm:"column:memo;size:64;" json:"memo" form:"memo"`                 // 备注
}

// 表名
func (Role) TableName() string {
	return TableName("role")
}

// 添加前
func (m *Role) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *Role) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 删除角色及关联数据
func (Role) Delete(ids []uint64) error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("id in (?)", ids).Delete(&Role{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("role_id in (?)", ids).Delete(&RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
