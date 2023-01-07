package model

import (
	"berpar/common"

	"gorm.io/gorm"
)


type SongList struct {
	ID uint `gorm:"primarykey" json:"id"`
	Title string `gorm:"type:varchar(255)" json:"title"`
	Pic string `gorm:"type:varchar(255)" json:"pic"`
	Introduction string `gorm:"type:text" json:"introduction"`
	Style string `gorm:"type:varchar(10)" json:"style"`
}

func AddSongList(songList SongList)*common.ResponseBody{
	err:=Db.Create(&songList).Error
	if err!=nil {
		return nil
	}
	return common.SuccessMessage("成功添加歌单！",songList.Title)
}

func DeleteSongList(id uint)*common.ResponseBody{
	err:=Db.Delete(&SongList{},id).Error
	if err!=nil {
		return nil
	}
	return common.SuccessMessage("成功删除歌单！",id)
}

func GetAllSongList() *common.ResponseBody{
	var songLists []SongList
	err:=Db.Model(&SongList{}).Where(&SongList{}).Find(&songLists).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("",songLists)
}


func UpdateSongListMsg(id int,title string)*common.ResponseBody{
	err:=Db.Model(&SongList{ID: uint(id)}).
	Select("title").
	Updates(map[string]interface{}{"title": title}).
	Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新歌单信息！",title)
}

func UpdateSongListPic(id int,pic string)*common.ResponseBody{
	err:=Db.Model(&SongList{ID: uint(id)}).
	Select("pic").
	Updates(map[string]interface{}{"pic": pic}).
	Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新歌单图片！",pic)
}

func GetSongListById(id int) *SongList{
	var songList SongList
	err:=Db.Model(&SongList{}).
	Where(&SongList{ID: uint(id)}).
	Find(&songList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return &songList
}

func SongListOfStyle(style string) *common.ResponseBody{
	var songLists []SongList
	err:=Db.Model(&SongList{}).Where("style like ?","%"+style+"%").Find(&songLists).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("指定类型歌单全集",songLists)
}

func SongListOfLikeTitle(title string) *common.ResponseBody{
	var songLists []SongList
	err:=Db.Model(&SongList{}).Where("title like ?","%"+title+"%").Find(&songLists).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("包含标题文字的歌单全集",songLists)
}