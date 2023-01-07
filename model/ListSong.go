package model

import (
	"berpar/common"

	"gorm.io/gorm"
)

type ListSong struct {
	ID uint `gorm:"primarykey" json:"id"`
	SongID uint `gorm:"type:int" json:"songId"`
	SongListID uint `gorm:"type:int" json:"songListId"`
}

func AddListSong(songId,songListId uint)*common.ResponseBody{
	listSong:=ListSong{SongID: songId,SongListID: songListId}
	err:=Db.Create(&listSong).Error
	if err!=nil {
		return nil
	}
	return common.SuccessMessage("成功添加歌单歌曲！",listSong)
}


func DeleteListSong(songId,songListId uint)*common.ResponseBody{
	err:=Db.Where(&ListSong{SongID: songId,SongListID: songListId}).Delete(&ListSong{}).Error
	if err!=nil {
		return nil
	}
	return common.SuccessMessage("成功删除歌单歌曲！",songId)
}

func ListSongOfId(id uint) *common.ResponseBody{
	var  listSongs []ListSong
	err:=Db.Model(&ListSong{}).
	Where(&ListSong{SongListID: id}).
	Find(&listSongs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("对应歌单id的所有歌曲",listSongs)
}

func UpdateListSongMsg(id,songId,songListId int)*common.ResponseBody{
	err:=Db.Model(&ListSong{ID: uint(id)}).
	Updates(map[string]interface{}{"song_id": songId,"song_list_id":songListId}).
	Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新歌单歌曲信息！",id)
}