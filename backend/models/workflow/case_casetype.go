package workflow

import (
	"time"

	"iris-ticket/backend/models/basemodel"

	"github.com/jinzhu/gorm"
)

// 工作流-工作流类型
type CaseCaseType struct {
	basemodel.Model
	CaseID     uint64 `gorm:"column:case_id;unique_index:uk_case_casetype_case_id;not null;"`         // 工作流ID
	CaseTypeID uint64 `gorm:"column:casetype_id;unique_index:uk_case_casetype_casetype_id;not null;"` // 工作流类型ID
}

// 表名
func (CaseCaseType) TableName() string {
	return TableName("case_casetype")
}

// 添加前
func (m *CaseCaseType) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *CaseCaseType) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 分配用户角色
// func (UserRole) SetRole(userid uint64, roleids []uint64) error {
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
// 	if err := tx.Where(&UserRole{UserID: userid}).Delete(&UserRole{}).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	if len(roleids) > 0 {
// 		for _, rid := range roleids {
// 			rm := new(UserRole)
// 			rm.RoleID = rid
// 			rm.UserID = userid
// 			if err := tx.Create(rm).Error; err != nil {
// 				tx.Rollback()
// 				return err
// 			}
// 		}
// 	}
// 	return tx.Commit().Error
// }
