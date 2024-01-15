package models

import "time"

// Auth2Collects 自定义第三张表  记录用户什么时候收藏文章
type Auth2Collects struct {
	UserID       int          `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"primaryKey:UserID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreateAt     time.Time
}
