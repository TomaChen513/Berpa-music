package model

import (
	"berpar/common"
	"time"

	"gorm.io/gorm"
)

type Singer struct {
	ID uint `gorm:"primarykey" json:"id"`
	Name string `gorm:"type:varchar(45)" json:"name"`
	Sex int `gorm:"type:tinyint" json:"sex"`
	Pic string `gorm:"type:varchar(255)" json:"pic"`
	Birth time.Time `gorm:"type:datetime" json:"birth"`
	Location string `gorm:"type:varchar(45)" json:"location"`
	Introduction string `gorm:"type:varchar(255)" json:"introduction"`
}


//  singer
func GetSinger() *common.ResponseBody{
	var allSinger []Singer
	err:=Db.Model(&Singer{}).
	Where(&Singer{}).
	Find(&allSinger).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("所有歌手",allSinger)
}

func GetSingerById(id int) *Singer{
	var singer Singer
	err:=Db.Model(&Singer{}).
	Where(&Singer{ID: uint(id)}).
	Find(&singer).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return &singer
}

func UpdateSingerPic(id int,pic string)*common.ResponseBody{
	err:=Db.Model(&Singer{ID: uint(id)}).
	Select("pic").
	Updates(map[string]interface{}{"pic": pic}).
	Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新歌手图片！",pic)
}

func AddSinger(singer Singer) *common.ResponseBody {
	// singer.Birth=time.Now()
	err:=Db.Create(&singer).Error
	if err!=nil {
		return nil
	}
	return common.SuccessMessage("成功添加歌手！",singer)
}

func DeleteSinger(id uint) *common.ResponseBody{
	err:=Db.Delete(&Singer{},id).Error
	if err!=nil {
		return nil
	}
	return common.SuccessMessage("删除成功歌手！",id)
}

func SingerOfName(name string) *common.ResponseBody{
	var singerList []Singer
	err:=Db.Model(&Singer{}).Where("name like ?","%"+name+"%").Find(&singerList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("歌手模糊查询",singerList)
}

func SingerOfSex(sex int) *common.ResponseBody{
	var singerList []Singer
	err:=Db.Model(&Singer{}).Where(&Singer{Sex: sex}).Find(&singerList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("歌手性别",singerList)
}

func UpdateSingerMsg(id int,location string)*common.ResponseBody{
	err:=Db.Model(&Singer{ID: uint(id)}).
	Select("location").
	Updates(map[string]interface{}{"location": location}).
	Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新歌手信息！",location)
}