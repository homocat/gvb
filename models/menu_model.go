package models

import (
	"main/models/ctype"
)

type MenuModel struct {
	MODEL

	MenuTitle    string        `gorm:"size:32" json:"menu_title"`
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`
	Slogan       string        `gorm:"size:64" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`
	AbstractTime int           `json:"abstract_time"` // 简介切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:BannerID:JoinReferences:ImageID" json:"banners"`
	BannerTime   int           `json:"banner_time"`         // 菜单图片切换时间, 0 表示不切换
	Sort         int           `gorm:"size:10" json:"sort"` // 菜单顺序
}
