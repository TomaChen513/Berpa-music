package controller

import (
	"berpar/model"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary  根据歌曲id获得歌曲信息
// @Produce json
// @Param id query string true "歌曲id"
// @Tags  song
// @Router /api/song/detail [get]
func SelectBySongId(ctx *gin.Context){
	id,_:=strconv.ParseInt(ctx.Query("id"),10,64)
	ctx.JSON(200,model.SelectBySongId(uint(id)))
}

// @Summary  根据歌手id获得歌曲信息
// @Produce json
// @Param singerId query string true "歌手id"
// @Tags  song
// @Router /api/song/singer/detail [get]
func SelectBySingerId(ctx *gin.Context){
	singerId,_:=strconv.ParseInt(ctx.Query("singerId"),10,64)
	ctx.JSON(200,model.SelectBySingerId(uint(singerId)))
}

// @Summary  根据歌手名称获得歌曲信息
// @Produce json
// @Param name query string true "歌手姓名"
// @Tags  song
// @Router /api/song/singerName/detail [get]
func SelectBySingerName(ctx *gin.Context){
	name:=ctx.Query("name")
	ctx.JSON(200,model.SelectBySingerName(name))
}

// @Summary  获得所有歌曲
// @Produce json
// @Tags  song
// @Router /api/song [get]
func GetAllSong(ctx *gin.Context){
	ctx.JSON(200,model.GetAllSong())
}



// @Summary  添加歌曲
// @Produce json
// @Tags  song
// @Router /api/song/add [post]
func AddSong(ctx *gin.Context){
	file, _:=ctx.FormFile("file")
	singerId,_:=strconv.Atoi(ctx.PostForm("id"))
	singer:=model.GetSingerById(singerId)
	// 写入对应文件
	path:="./song/" + strconv.Itoa(time.Now().Nanosecond())+singer.Name+".mp3"
	ctx.SaveUploadedFile(file, path)
	ctx.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// 将文件名映射到singer数据库中
	ctx.JSON(200,model.AddSong(singerId,path[1:],file.Filename))
}

// @Summary  删除歌曲
// @Produce json
// @Tags  song
// @Router /api/song/delete [get]
func DeleteSong(ctx *gin.Context){
	id,_:=strconv.Atoi(ctx.Query("id"))
	songUrl:=model.GetSongUrlById(uint(id))
    os.Remove("."+songUrl)
	ctx.JSON(200,model.DeleteSong(uint(id)))
}

// @Summary  更新歌曲信息
// @Produce json
// @Tags  song
// @Router /api/song/update [post]
func UpdateSongMsg(ctx *gin.Context){
	name:=ctx.PostForm("name")
	id,_:=strconv.Atoi(ctx.PostForm("id"))
	ctx.JSON(200,model.UpdateSongMsg(id,name))
}

// @Summary  更新歌曲图片
// @Produce json
// @Tags  song
// @Router /api/song/img/update [post]
func UpdateSongPic(ctx *gin.Context){
	file, _:=ctx.FormFile("file")
	songId,_:=strconv.Atoi(ctx.PostForm("id"))

	song:=model.GetSongById(songId)
	// 写入对应文件
	path:="./img/songPic/" + strconv.Itoa(time.Now().Nanosecond())+song.Name+".jpg"
	ctx.SaveUploadedFile(file, path)
	ctx.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// 将文件名映射到singer数据库中
	ctx.JSON(200,model.UpdateSongPic(songId,path[1:]))
}

// @Summary  更新歌曲
// @Produce json
// @Tags  song
// @Router /api/song/url/update [post]
func UpdateSongUrl(ctx *gin.Context){
	file, _:=ctx.FormFile("file")
	songId,_:=strconv.Atoi(ctx.PostForm("id"))
	song:=model.GetSongById(songId)
	preUrl:=song.URL
	os.Remove("."+preUrl)
	// 写入对应文件
	path:="./song/" + strconv.Itoa(time.Now().Nanosecond())+".mp3"
	ctx.SaveUploadedFile(file, path)
	ctx.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// 将文件名映射到singer数据库中
	ctx.JSON(200,model.UpdateSongUrl(songId,path[1:]))
}

func ModifySongUrl(ctx *gin.Context){
	ctx.JSON(200,model.ModifySongUrl())
}