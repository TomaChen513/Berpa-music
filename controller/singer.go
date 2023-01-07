package controller

import (
	"berpar/model"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary  返回所有歌手
// @Produce json
// @Tags  singer
// @Router /api/singer [get]
func GetSinger(ctx *gin.Context){
	ctx.JSON(200,model.GetSinger())
}

// @Summary  更新歌手图片
// @Produce json
// @Tags  singer
// @Router /api/singer/avatar/update [post]
func UpdateSingerPic(ctx *gin.Context){
	file, _:=ctx.FormFile("file")
	singerId,_:=strconv.Atoi(ctx.PostForm("id"))
	singer:=model.GetSingerById(singerId)
	// 写入对应文件
	path:="./img/singerPic/" + strconv.Itoa(time.Now().Nanosecond())+singer.Name+".jpg"
	ctx.SaveUploadedFile(file, path)
	ctx.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// 将文件名映射到singer数据库中
	ctx.JSON(200,model.UpdateSingerPic(singerId,path[1:]))
}

// @Summary  添加歌手
// @Produce json
// @Tags  singer
// @Router /api/singer/add [post]
func AddSinger(ctx *gin.Context){
	name:=ctx.PostForm("name")
	sex:=ctx.PostForm("sex")
	location:=ctx.PostForm("location")
	pic:="testpic"
	introduction:=ctx.PostForm("introduction")
	singer:=model.Singer{}
	singer.Name=name
	singer.Sex,_=strconv.Atoi(sex)
	singer.Location=location
	singer.Pic=pic
	singer.Introduction=introduction
	currTime,_:=time.Parse("2006-01-02 15:04:05", "1969-08-08 17:58:31")
	singer.Birth=currTime
	ctx.JSON(200,model.AddSinger(singer))
}

// @Summary  删除歌手
// @Produce json
// @Tags  singer
// @Router /api/singer/delete [get]
func DeleteSinger(ctx *gin.Context){
	id,_:=strconv.Atoi(ctx.Query("id"))
	ctx.JSON(200,model.DeleteSinger(uint(id)))
}

// @Summary  根据歌手名查找歌手
// @Produce json
// @Tags  singer
// @Router /api/singer/name/detail [get]
func SingerOfName(ctx *gin.Context){
	name:=ctx.Query("name")
	ctx.JSON(200,model.SingerOfName(name))
}

// @Summary  根据歌手性别查找歌手
// @Produce json
// @Tags  singer
// @Router /api/singer/sex/detail [get]
func SingerOfSex(ctx *gin.Context){
	sex,_:=strconv.Atoi(ctx.Query("sex"))
	ctx.JSON(200,model.SingerOfSex(sex))
}

// @Summary  更新歌手信息
// @Produce json
// @Tags  singer
// @Router /api/singer/update [post]
func UpdateSingerMsg(ctx *gin.Context){
	location:=ctx.PostForm("location")
	id,_:=strconv.Atoi(ctx.PostForm("id"))
	ctx.JSON(200,model.UpdateSingerMsg(id,location))
}

