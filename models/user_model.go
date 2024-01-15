package models

import (
	"main/models/ctype"
)

type UserModel struct {
	// gorm.Model 包含删除逻辑, 需要再改
	MODEL
	NickName string `gorm:"size:36" json:"nickname"`  // 昵称
	UserName string `gorm:"size:36" json:"username"`  // 账号
	Password string `gorm:"size:128" json:"password"` // 密码

	Avatar         string           `gorm:"size:256" json:"avatar_id"`            // 头像id
	Email          string           `gorm:"size:128" json:"email"`                // 邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                   // 电话
	Token          string           `gorm:"size:64" json:"token"`                 // 其他平台唯一id
	IP             string           `gorm:"size:20" json:"ip"`                    // ip地址
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`         // 权限
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)j" json:"sign_status"` // 注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:AuthID" json:"-"`
	CollectsModels []ArticleModel   `gorm:"many2many:auth2_collects;joinForeignKey:AuthID:JoinReferences:ArticleID" json:"-"`
}
