package model

import "gorm.io/gorm"


type Category struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	Name     string    `gorm:"type:varchar(255);not null;unique;comment:分类名" json:"name"`
	Mid      *uint     `gorm:"type:int;comment:菜单子项ID" json:"mid"`
	Homeshow bool      `gorm:"type:bool;comment:主页是否显示;not null;" json:"homeshow,omitempty"`
}

type GetCategoryTy struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Mid      uint   `json:"mid"`
	Homeshow bool   `json:"homeshow"`
}

// GetCategory 获取分类
func GetCategory(homeshow bool) []Category {
	var data []Category
	err := Db.Model(&Category{}).
		Where(&Category{Homeshow: homeshow}).
		Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return data
}