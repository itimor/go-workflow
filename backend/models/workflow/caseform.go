package workflow

import (
	"go-workflow/backend/models/db"
	"time"

	"go-workflow/backend/models/basemodel"

	"github.com/jinzhu/gorm"
)

// 发布表单
type CaseForm struct {
	basemodel.Model
	Name string `gorm:"column:name;size:32;not null;" json:"name" form:"name"` // 名称
	Code string `gorm:"column:code;size:32;not null;" json:"code" form:"code"` // 代码
}

// 表名
func (CaseForm) TableName() string {
	return TableName("caseform")
}

// 添加前
func (m *CaseForm) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *CaseForm) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 删除
func (CaseForm) Delete(ids []uint64) error {
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
	if err := tx.Where("id in (?)", ids).Delete(&CaseForm{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
