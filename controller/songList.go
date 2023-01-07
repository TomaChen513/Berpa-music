package controller

import (
	"berpar/model"
	"fmt"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

// @Summary  添加歌单
// @Produce json
// @Tags  songList
// @Router /api/songList/add [post]
func AddSongList(ctx *gin.Context){
	title:=ctx.PostForm("title")
	songList:=model.SongList{Title: title}
	ctx.JSON(200,model.AddSongList(songList))
}


// @Summary  删除歌单
// @Produce json
// @Tags  songList
// @Router /api/songList/delete [get]
func DeleteSongList(ctx *gin.Context){
	id,_:=strconv.Atoi(ctx.Query("id"))
	ctx.JSON(200,model.DeleteSongList(uint(id)))
}

// @Summary  获得所有歌单
// @Produce json
// @Tags  songList
// @Router /api/songList [get]
func GetAllSongList(ctx *gin.Context){
	ctx.JSON(200,model.GetAllSongList())
}


// @Summary  更新歌单信息
// @Produce json
// @Tags  songList
// @Router /api/songList/update [post]
func UpdateSongListMsg(ctx *gin.Context){
	name:=ctx.PostForm("title")
	id,_:=strconv.Atoi(ctx.PostForm("id"))
	ctx.JSON(200,model.UpdateSongListMsg(id,name))
}

// @Summary  更新歌单图片
// @Produce json
// @Tags  songList
// @Router /api/songList/img/update [post]
func UpdateSongListPic(ctx *gin.Context){
	file, _:=ctx.FormFile("file")
	songListId,_:=strconv.Atoi(ctx.PostForm("id"))
	songList:=model.GetSongListById(songListId)
	// 写入对应文件
	path:="./img/songListPic/" + strconv.Itoa(time.Now().Nanosecond())+songList.Title+".jpg"
	ctx.SaveUploadedFile(file, path)
	ctx.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// 将文件名映射到singer数据库中
	ctx.JSON(200,model.UpdateSongListPic(songListId,path[1:]))
}

// @Summary  根据类型获得歌单
// @Produce json
// @Tags  songList
// @Router /api/songList/style/detail [get]
func SongListOfStyle(ctx *gin.Context){
	style:=ctx.Query("style")
	ctx.JSON(200,model.SongListOfStyle(style))
}

// @Summary  返回标题包含文字的歌单
// @Produce json
// @Tags  songList
// @Router /api/songList/likeTitle/detail [get]
func SongListOfLikeTitle(ctx *gin.Context){
	title:=ctx.Query("title")
	ctx.JSON(200,model.SongListOfLikeTitle(title))
}