package model

import (
	"berpar/common"
	"time"

	"gorm.io/gorm"
)

type Collect struct {
	ID uint `gorm:"primarykey" json:"id"`
	UserID uint `gorm:"type:int" json:"userId"`
	Type int `gorm:"type:tinyint" json:"type"`
	SongID uint `gorm:"type:int" json:"songId"`
	SongListID uint `gorm:"type:int" json:"songListId"`
	CreateTime time.Time `gorm:"type:datetime" json:"createTime"`
}



func AddCollection(collect Collect)*common.ResponseBody{
	err:=Db.Create(&collect).Error
	if err!=nil {
		return nil
	}
	return common.SuccessMessage("成功添加收藏！",collect)
}


func DeleteCollect(userId,songId uint) *common.ResponseBody{
	err:=Db.Where("user_id=?",userId).Where("song_id=?",songId).Delete(&Collect{}).Error
	if err!=nil {
		return nil
	}
	return common.SuccessMessage("成功取消收藏！",userId)
}

func IsCollection(userId,songId int) *common.ResponseBody{
	var collect Collect
	err:=Db.Model(&Collect{}).
	Where(&Collect{UserID: uint(userId),SongID: uint(songId)}).
	Find(&collect).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return common.SuccessMessage("未收藏！",userId)
	}
	return common.SuccessMessage("已收藏！",userId)
}

func CollectionOfUser(userId int) *common.ResponseBody{
	var collects []Collect
	err:=Db.Model(&Collect{}).
	Where(&Collect{UserID: uint(userId)}).
	Find(&collects).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("收藏列表！",collects)
}

