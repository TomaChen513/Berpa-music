package model

import (
	"time"

	"gorm.io/gorm"

	"berpar/common"
)

type Song struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	SingerID     uint      `gorm:"type:int" json:"singerId"`
	Name         string    `gorm:"type:varchar(45)" json:"name"`
	Introduction string    `gorm:"type:varchar(255)" json:"introduction"`
	CreateTime   time.Time `gorm:"type:datetime" json:"createTime"`
	UpdateTime   time.Time `gorm:"type:dateTime" json:"updateTime"`
	Pic          string    `gorm:"type:varchar(255)" json:"pic"`
	Lyric        string    `gorm:"type:text" json:"lyric"`
	URL          string    `gorm:"type:varchar(255)" json:"url"`
}

// /song

// /detail
func SelectBySongId(id uint) *common.ResponseBody {
	var song []Song
	err := Db.Model(&Song{}).
		Where(&Song{ID: id}).
		Find(&song).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	return common.SuccessMessage("歌曲id", song)
}

func GetSongUrlById(id uint) string {
	var song Song
	err := Db.Model(&Song{}).
		Where(&Song{ID: id}).
		Find(&song).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return ""
	}
	return song.URL
}

// /singer/detail
func SelectBySingerId(id uint) *common.ResponseBody {
	var songList []Song
	err := Db.Model(&Song{}).Where(&Song{SingerID: id}).Find(&songList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("歌手id", songList)
}

// /singerName/detail
func SelectBySingerName(name string) *common.ResponseBody {
	var songList []Song
	err := Db.Model(&Song{}).Where("name like ?", "%"+name+"%").Find(&songList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("歌曲全集", songList)
}

// /
func GetAllSong() *common.ResponseBody {
	var songList []Song
	err := Db.Model(&Song{}).Where(&Song{}).Find(&songList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("", songList)
}

func AddSong(singerId int, url, name string) *common.ResponseBody {
	song := Song{SingerID: uint(singerId), URL: url}
	song.CreateTime = time.Now()
	song.UpdateTime = time.Now()
	song.Name = name
	err := Db.Create(&song).Error
	if err != nil {
		return nil
	}
	return common.SuccessMessage("成功添加歌曲！", url)
}

func DeleteSong(id uint) *common.ResponseBody {
	err := Db.Delete(&Song{}, id).Error
	if err != nil {
		return nil
	}
	return common.SuccessMessage("成功删除歌曲！", id)
}

func UpdateSongMsg(id int, name string) *common.ResponseBody {
	err := Db.Model(&Song{ID: uint(id)}).
		Select("name").
		Updates(map[string]interface{}{"name": name}).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新歌曲信息！", name)
}

func GetSongById(id int) *Song {
	var song Song
	err := Db.Model(&Song{}).
		Where(&Song{ID: uint(id)}).
		Find(&song).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return &song
}

func UpdateSongPic(id int, pic string) *common.ResponseBody {
	err := Db.Model(&Song{ID: uint(id)}).
		Select("pic").
		Updates(map[string]interface{}{"pic": pic}).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新歌曲图片！", pic)
}

func UpdateSongUrl(songId int, path string) *common.ResponseBody {
	err := Db.Model(&Song{ID: uint(songId)}).
		Select("url").
		Updates(map[string]interface{}{"url": path}).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新歌曲！", path)
}

func ModifySongUrl() *common.ResponseBody {
	var songList []Song
	err := Db.Model(&Song{}).Where(&Song{}).Find(&songList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	// songList[1].URL = songList[1].URL[:5] + "s" + songList[1].URL[5:]
	for i := 0; i < len(songList); i++ {
		oldUrl := songList[i].URL
		newUrl := oldUrl[:5] + "s" + oldUrl[5:]
		UpdateSongUrl(int(songList[i].ID),newUrl)
	}
	return common.SuccessMessage("modify", songList)

}
