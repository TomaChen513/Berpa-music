package model

import (
	"berpar/common"

	"gorm.io/gorm"
)

type RankList struct {
	ID         uint `gorm:"primarykey" json:"id"`
	SongListID uint `gorm:"column:songListId;type:int" json:"songListId"`
	ConsumerID uint `gorm:"column:consumerId;type:int" json:"consumerId"`
	Score      uint `gorm:"type:int = '0'" json:"score"`
}

// rankList
func GetScoreBySongList(songListId uint) *common.ResponseBody {
	var rankList []RankList
	err := Db.Model(&RankList{}).
		Where(&RankList{SongListID: songListId}).
		Find(&rankList).Error
	if err != nil {
		return nil
	}
	var sum float64
	if len(rankList) != 0 {
		sum = 0.0
		for _, list := range rankList {
			sum += float64(list.Score)
		}
		sum /= float64(len(rankList))
	}
	return common.SuccessMessage("评分为：", sum)
}

// /user
func GetScoreByUser(songListId, consumerId uint) *common.ResponseBody {
	var rankList RankList
	err := Db.Model(&RankList{}).
		Where(&RankList{SongListID: songListId, ConsumerID: consumerId}).
		Find(&rankList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("评分为", rankList.Score)
}

// /add
func AddScore(data *RankList) *common.ResponseBody {
	Db.Create(&data)
	return common.SuccessMessage("打分成功", nil)
}
