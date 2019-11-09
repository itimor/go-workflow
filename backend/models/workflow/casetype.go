package workflow

import (
	"go-workflow/backend/models/db"
	"time"

	"go-workflow/backend/models/basemodel"

	"github.com/jinzhu/gorm"
)

// 工作流类型
type CaseType struct {
	basemodel.Model
	Name   string `gorm:"column:name;size:32;not null;" json:"name" form:"name"`               // 名称
	Status uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"` // 状态(1:启用 2:禁用)
	Form   uint8  `gorm:"column:form;type:tinyint(1);not null;" json:"form" form:"form"`       // 工作流类型表单(1:发布 2:请假 3...)
	Memo   string `gorm:"column:memo;size:64;" json:"memo" form:"memo"`                        // 备注
}

// 表名
func (CaseType) TableName() string {
	return TableName("casetype")
}

// 添加前
func (m *CaseType) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *CaseType) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 删除
func (CaseType) Delete(ids []uint64) error {
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
	if err := tx.Where("id in (?)", ids).Delete(&CaseType{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("casetype_id in (?)", ids).Delete(&CaseTypeStep{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
