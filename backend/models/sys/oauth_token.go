package sys

import (
	"time"

	"iris-ticket/backend/models/basemodel"
	"iris-ticket/backend/models/db"

	"github.com/jinzhu/gorm"
)

type OauthToken struct {
	basemodel.Model
	Token     string `gorm:"not null default '' comment('Token') VARCHAR(191)" json:"access_token"`
	UserId    uint64 `gorm:"not null default '' comment('UserId') VARCHAR(191)"`
	Secret    string `gorm:"not null default '' comment('Secret') VARCHAR(191)"`
	ExpressIn int64  `gorm:"not null default 0 comment('是否是标准库') BIGINT(20)"`
	Revoked   bool
}

// 表名
func (OauthToken) TableName() string {
	return TableName("oauth_token")
}

// 添加前
func (m *OauthToken) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *OauthToken) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 删除角色及关联数据
func (OauthToken) Delete(roleids []uint64) error {
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
	if err := tx.Where("id in (?)", roleids).Delete(&OauthToken{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
