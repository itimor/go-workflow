package workflow

import (
	"go-workflow/backend/models/basemodel"
	"go-workflow/backend/models/db"
	
	"time"

	"github.com/jinzhu/gorm"
)

// 工作流操作
type CaseOpera struct {
	basemodel.Model
	UserID  uint64 `gorm:"column:user_id;unique_index:uk_caseopera_user_id;not null;"`          // 执行人ID
	CaseID  uint64 `gorm:"column:case_id;unique_index:uk_caseopera_case_id;not null;"`          // 工作流类型ID
	Content string `gorm:"column:content;type:text;" json:"content" form:"content"`             // 回复
	Status  uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"` // 状态(1:待提交 2:审核中 3:审核驳回 4:执行中 5:执行驳回 6:执行完成 7:完成关闭 8:驳回关闭 9:撤销关闭 10:重新编辑 11:回复)
}

// 表名
func (CaseOpera) TableName() string {
	return TableName("caseopera")
}

// 添加前
func (m *CaseOpera) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *CaseOpera) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 删除
func (CaseOpera) Delete(ids []uint64) error {
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
	if err := tx.Where("id in (?)", ids).Delete(&CaseOpera{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
