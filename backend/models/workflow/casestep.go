package workflow

import (
	"time"

	"iris-ticket/backend/models/basemodel"

	"github.com/jinzhu/gorm"
)

// 工作流步骤
type CaseStep struct {
	basemodel.Model
	UserID uint64 `gorm:"column:user_id;unique_index:uk_casestep_user_id;not null;"`           // 执行人ID
	CaseID uint64 `gorm:"column:casetype_id;unique_index:uk_casestep_case_id;not null;"`       // 工作流类型ID
	Step   uint8  `gorm:"column:step;type:tinyint(1);not null;" json:"step" form:"step"`       // 执行步骤(1 2 3 4)
	Status uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"` // 状态(1:未执行 2:已执行)
}

// 表名
func (CaseStep) TableName() string {
	return TableName("casestep")
}

// 添加前
func (m *CaseStep) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *CaseStep) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 删除角色及关联数据
// func (CaseType) Delete(roleids []uint64) error {
// 	tx := db.DB.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	if err := tx.Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	if err := tx.Where("id in (?)", roleids).Delete(&Role{}).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	return tx.Commit().Error
// }
