package controller

import (
	"berpar/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary  给歌单添加歌曲
// @Produce json
// @Tags  listSong
// @Router /api/listSong/add [post]
func AddListSong(ctx *gin.Context){
	songId,_:=strconv.ParseInt(ctx.PostForm("songId"),10,64)
	songListId,_:=strconv.ParseInt(ctx.PostForm("songListId"),10,64)
	ctx.JSON(200,model.AddListSong(uint(songId),uint(songListId)))
}


// @Summary  删除歌单里的歌曲
// @Produce json
// @Tags  listSong
// @Router /api/listSong/delete [get]
func DeleteListSong(ctx *gin.Context){
	songId,_:=strconv.ParseInt(ctx.Query("songId"),10,64)
	songListId,_:=strconv.ParseInt(ctx.Query("songListId"),10,64)
	ctx.JSON(200,model.DeleteListSong(uint(songId),uint(songListId)))
}



// @Summary  根据歌曲id获得歌单歌曲信息
// @Produce json
// @Tags  listSong
// @Router /api/listSong/detail [get]
func ListSongOfId(ctx *gin.Context){
	id,_:=strconv.ParseInt(ctx.Query("songListId"),10,64)
	ctx.JSON(200,model.ListSongOfId(uint(id)))
}

// @Summary  更新歌单的歌曲信息
// @Produce json
// @Tags  listSong
// @Router /api/listSong/update [post]
func UpdateListSongMsg(ctx *gin.Context){
	id,_:=strconv.Atoi(ctx.PostForm("id"))
	songId,_:=strconv.Atoi(ctx.PostForm("songId"))
	songListId,_:=strconv.Atoi(ctx.PostForm("songListId"))
	ctx.JSON(200,model.UpdateListSongMsg(id,songId,songListId))
}