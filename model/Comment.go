package model

import (
	"berpar/common"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID uint `gorm:"primarykey" json:"id"`
	UserId uint `gorm:"type:int" json:"userId"`
	SongId uint `gorm:"type:int" json:"songId"`
	SongListId uint `gorm:"type:int" json:"songListId"`
	Content string `gorm:"type:varchar(255)" json:"content"`
	CreateTime time.Time `gorm:"type:datetime" json:"createTime"`
	Type int `gorm:"type:tinyint" json:"type"`
	Up uint `gorm:"type:int;default 0" json:"up"`
}

func AddComment(comment Comment)*common.ResponseBody{
	comment.CreateTime=time.Now()
	err:=Db.Create(&comment).Error
	if err!=nil {
		return nil
	}
	return common.SuccessMessage("成功添加评论！",comment)
}

func DeleteComment(id uint) *common.ResponseBody{
	err:=Db.Delete(&Comment{},id).Error
	if err!=nil {
		return nil
	}
	return common.SuccessMessage("成功删除评论！",id)
}

func CommentOfSongId(id uint) *common.ResponseBody{
	var comments []Comment
	err:=Db.Model(&Comment{}).Where(&Comment{SongId: id}).Find(&comments).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("指定歌曲评论",comments)
}

func CommentOfSongListId(id uint) *common.ResponseBody{
	var comments []Comment
	err:=Db.Model(&Comment{}).Where(&Comment{SongListId: id}).Find(&comments).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("指定歌单评论",comments)
}

func UpdateCommentMsg(id,up int)*common.ResponseBody{
	err:=Db.Model(&Comment{ID: uint(id)}).
	Select("up").
	Updates(map[string]interface{}{"up": up}).
	Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("点赞成功！",up)
}