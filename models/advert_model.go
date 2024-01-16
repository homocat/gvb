package models

// AdvertModel 广告表
type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:32" json:"title"`
	Href   string `json:"href"` // 跳转连接
	Images string `json:"images"`
	IsShow bool   `json:"isShow"` // 是否展示
}
