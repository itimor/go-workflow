package workflow

import (
	"go-workflow/backend/models/db"
	"time"

	"iris-ticket/backend/models/basemodel"

	"github.com/jinzhu/gorm"
)

// 工作流类型步骤
type CaseTypeStep struct {
	basemodel.Model
	Name       string `gorm:"column:name;size:32;not null;" json:"name" form:"name"`                 // 节点名称
	UserID     uint64 `gorm:"column:user_id;unique_index:uk_casetypestep_user_id;not null;"`         // 执行人ID
	CaseTypeID uint64 `gorm:"column:casetype_id;unique_index:uk_casetypestep_casetype_id;not null;"` // 工作流类型ID
	Step       uint8  `gorm:"column:step;type:tinyint(1);not null;" json:"step" form:"step"`         // 执行步骤(1 2 3 4)
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
