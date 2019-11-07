package workflow

import (
	"time"

	"iris-ticket/backend/models/basemodel"
	"iris-ticket/backend/models/sys"

	"github.com/jinzhu/gorm"
)

// 工作流
type Case struct {
	basemodel.Model
	Name       string   `gorm:"column:name;size:32;not null;" json:"name" form:"name"`               // 名称
	CreateUser sys.User `gorm:"foreignkey:UserID"`                                                   // 创建人
	Content    string   `gorm:"column:content;type:text;" json:"content" form:"content"`             // 工单需求
	Result     string   `gorm:"column:result;type:text;" json:"result" form:"result"`                // 工单结果
	Status     uint8    `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"` // 状态(1:待提交 2:审核中 3:审核驳回 4:执行中 5:执行驳回 6:执行完成 7:完成关闭 8:驳回关闭 9:撤销关闭)
	Step       uint8    `gorm:"column:step;type:tinyint(1);not null;" json:"step" form:"step"`       // 当前流程步骤(1 2 3 4)
	Memo       string   `gorm:"column:memo;size:64;" json:"memo" form:"memo"`                        // 备注
}

// 表名
func (Case) TableName() string {
	return TableName("case")
}

// 添加前
func (m *Case) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *Case) BeforeUpdate(scope *gorm.Scope) error {
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
