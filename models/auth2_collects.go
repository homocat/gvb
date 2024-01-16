package models

import "time"

// Auth2Collects 记录用户收藏文章的自定义第三张表
type Auth2Collects struct {
	UserID       int          `gorm:"primaryKey"`           // 用户ID
	UserModel    UserModel    `gorm:"primaryKey:UserID"`    // 用户模型
	ArticleID    uint         `gorm:"primaryKey"`           // 文章ID
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"` // 文章模型
	CreateAt     time.Time    // 创建时间
}
