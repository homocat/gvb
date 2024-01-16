package models

type FeedBackModel struct {
	MODEL
	Email        string `gorm:"size:64" json:"email"`
	Content      string `gorm:"size:128" json:"content"`
	ApplyContent string `gorm:"size:128" json:"apply_content"` // 回复内容
	IsApply      bool   `json:"is_apply"`                      // 是否回复
}
