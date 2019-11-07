package workflow

import (
	"time"

	"iris-ticket/backend/models/basemodel"
	"iris-ticket/backend/models/sys"

	"github.com/jinzhu/gorm"
)

// 工作流类型
type CaseType struct {
	basemodel.Model
	Name       string   `gorm:"column:name;size:32;not null;" json:"name" form:"name"`               // 名称
	CreateUser sys.User `gorm:"foreignkey:UserID"`                                                   // 创建人
	Status     uint8    `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"` // 状态(1:启用 2:禁用)
	Form       uint8    `gorm:"column:form;type:tinyint(1);not null;" json:"form" form:"form"`       // 工作流类型表单(1:发布 2:请假 3...)
	Memo       string   `gorm:"column:memo;size:64;" json:"memo" form:"memo"`                        // 备注
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
