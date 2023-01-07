package controller

import (
	"berpar/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary  添加收藏
// @Produce json
// @Tags  collect
// @Router /api/collection/add [post]
func AddCollection(ctx *gin.Context){
	userId,_:=strconv.Atoi(ctx.PostForm("userId"))
	songId,_:=strconv.Atoi(ctx.PostForm("songId"))
	songListId,_:=strconv.Atoi(ctx.PostForm("songListId"))
	types,_:=strconv.Atoi(ctx.PostForm("type"))
	collection:=model.Collect{UserID: uint(userId),CreateTime: time.Now()}
	if types==0 {
		collection.SongID=uint(songId)
	}else{
		collection.SongListID=uint(songListId)
	}
	ctx.JSON(200,model.AddCollection(collection))
}

// @Summary  取消收藏
// @Produce json
// @Tags  collect
// @Router /api/collection/delete [post]
func DeleteCollect(ctx *gin.Context){
	userId,_:=strconv.Atoi(ctx.PostForm("userId"))
	songId,_:=strconv.Atoi(ctx.PostForm("songId"))
	ctx.JSON(200,model.DeleteCollect(uint(userId),uint(songId)))
}

// @Summary  是否收藏歌曲
// @Produce json
// @Tags  collect
// @Router /api/collection/status [post]
func IsCollection(ctx *gin.Context){
	userId,_:=strconv.Atoi(ctx.Query("userId"))
	songId,_:=strconv.Atoi(ctx.Query("songId"))
	ctx.JSON(200,model.IsCollection(userId,songId))
}

// @Summary  返回指定用户的收藏列表
// @Produce json
// @Tags  collect
// @Router /api/collection/detail [get]
func CollectionOfUser(ctx *gin.Context){
	userId,_:=strconv.Atoi(ctx.Query("userId"))
	ctx.JSON(200,model.CollectionOfUser(userId))
}