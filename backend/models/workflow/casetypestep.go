package workflow

import (
	"go-workflow/backend/models/basemodel"
	"go-workflow/backend/models/db"

	"time"

	"github.com/jinzhu/gorm"
)

// 工作流类型步骤
type CaseTypeStep struct {
	basemodel.Model
	Name       string `gorm:"column:name;size:32;not null;" json:"name" form:"name"`              // 名称
	Type       uint8    `gorm:"column:type;type:tinyint(1);not null;" json:"type" form:"type"`      // 状态(0:提交 1:审核 2:执行 3:关闭)
	UserID     uint64    `gorm:"column:user_id;not null;" json:"user_id" form:"user_id"`             // 执行人ID
	CaseTypeID uint64    `gorm:"column:casetype_id;not null;" json:"casetype_id" form:"casetype_id"` // 工作流类型ID
	Step       uint8    `gorm:"column:step;type:tinyint(1);not null;" json:"step" form:"step"`      // 执行步骤(1 2 3 4)
}

// 表名
func (CaseTypeStep) TableName() string {
	return TableName("casetypestep")
}

// 添加前
func (m *CaseTypeStep) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *CaseTypeStep) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 删除
func (CaseTypeStep) Delete(ids []uint64) error {
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
	if err := tx.Where("id in (?)", ids).Delete(&CaseTypeStep{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
