package workflowform

import (
	"go-workflow/backend/models/db"
	"time"

	"iris-ticket/backend/models/basemodel"
	"iris-ticket/backend/models/sys"

	"github.com/jinzhu/gorm"
)

// 发布表单
type Deploy struct {
	basemodel.Model
	Name       string   `gorm:"column:name;size:32;not null;" json:"name" form:"name"`   // 名称
	CreateUser sys.User `gorm:"foreignkey:UserID"`                                       // 创建人
	Content    string   `gorm:"column:content;type:text;" json:"content" form:"content"` // 发布说明
	Result     string   `gorm:"column:result;type:text;" json:"result" form:"result"`    // 发布结果
	Memo       string   `gorm:"column:memo;size:64;" json:"memo" form:"memo"`            // 备注
}

// 表名
func (Deploy) TableName() string {
	return TableName("deploy")
}

// 添加前
func (m *Deploy) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *Deploy) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 删除
func (Deploy) Delete(ids []uint64) error {
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
	if err := tx.Where("id in (?)", ids).Delete(&Deploy{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
