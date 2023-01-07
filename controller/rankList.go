package controller

import (
	"berpar/model"
	"strconv"

	"github.com/gin-gonic/gin"
)


// @Summary  获取指定歌单评分
// @Produce json
// @Tags  rankList
// @Router /api/rankList [get]
func GetScoreBySongList(ctx *gin.Context){
	songListId,_:=strconv.ParseUint(ctx.Query("songListId"),10,64)
	ctx.JSON(200,model.GetScoreBySongList(uint(songListId)))
}

// @Summary  根据用户获取指定歌单评分
// @Produce json
// @Tags  rankList
// @Router /api/rankList/user [get]
func GetScoreByUser(ctx *gin.Context){
	songListId,_:=strconv.ParseUint(ctx.Query("songListId"),10,64)
	consumerId,_:=strconv.ParseUint(ctx.Query("consumerId"),10,64)

	ctx.JSON(200,model.GetScoreByUser(uint(songListId),uint(consumerId)))
}


// @Summary 提交评分
// @Produce json
// @Tags  rankList
// @Router /api/rankList/add [post]
func AddScore(ctx *gin.Context){
	var data model.RankList	
	songListId,_:=strconv.ParseInt(ctx.PostForm("songListId"),10,64)
	consumerId,_:=strconv.ParseInt(ctx.PostForm("consumerId"),10,64)
	score,_:=strconv.ParseInt(ctx.PostForm("score"),10,64)
	data.SongListID=uint(songListId)
	data.ConsumerID=uint(consumerId)
	data.Score=uint(score)
	ctx.JSON(200,model.AddScore(&data))
}