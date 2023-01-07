package controller

import (
	"berpar/model"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"

)

// @Summary  用户注册
// @Produce json
// @Tags  user
// @Router /api/user/add [post]
func AddUser(ctx *gin.Context){
	userName:=ctx.PostForm("username")
	passWord:=ctx.PostForm("password")
	ctx.JSON(200,model.AddUser(userName,passWord))
}


// @Summary  登录判断
// @Produce json
// @Tags  user
// @Router /api/user/login/status [post]
func LoginStatus(ctx *gin.Context){
	userName:=ctx.PostForm("username")
	passWord:=ctx.PostForm("password")
	session:=sessions.Default(ctx)
	session.Set("username",userName)
	ctx.JSON(200,model.VerifyPasswd(userName,passWord))
}

// @Summary  获得所有用户
// @Produce json
// @Tags  user
// @Router /api/user [get]
func GetAllUser(ctx *gin.Context){
	ctx.JSON(200,model.GetAllUser())
}

// @Summary  根据用户id获得歌曲信息
// @Produce json
// @Tags  user
// @Router /api/user/detail [get]
func SelectByUserId(ctx *gin.Context){
	temp:=ctx.Query("id")
	var userId int64
	if strconv.Itoa(len(temp))=="0" {
		userId=0
	}else{
		userId,_=strconv.ParseInt(temp,10,64)
	}
	// ctx.String(200,strconv.Itoa(int(userId)))
	ctx.JSON(200,model.SelectByUserId(uint(userId)))
}

// @Summary  删除用户
// @Produce json
// @Tags  user
// @Router /api/user/delete [get]
func DeleteUser(ctx *gin.Context){
	id,_:=strconv.Atoi(ctx.Query("id"))
	ctx.JSON(200,model.DeleteUser(uint(id)))
}

// @Summary  更新用户信息
// @Produce json
// @Tags  user
// @Router /api/user/update [post]
func UpdateUserMsg(ctx *gin.Context){
	email:=ctx.PostForm("email")
	id,_:=strconv.Atoi(ctx.PostForm("id"))
	ctx.JSON(200,model.UpdateUserMsg(id,email))
}


// @Summary  更新用户密码
// @Produce json
// @Tags  user
// @Router /api/user/updatePassword [post]
func UpdateUserPassword(ctx *gin.Context){
	password:=ctx.PostForm("password")
	id,_:=strconv.Atoi(ctx.PostForm("id"))
	ctx.JSON(200,model.UpdateUserPassword(id,password))
}


// @Summary  更新用户图片
// @Produce json
// @Tags  user
// @Router /api/user/avator/update [post]
func UpdateUserPic(ctx *gin.Context){
	file, _:=ctx.FormFile("file")
	userId,_:=strconv.Atoi(ctx.PostForm("id"))

	// 写入对应文件
	path:="./img/avatorImages/" + strconv.Itoa(time.Now().Nanosecond())+".jpg"
	ctx.SaveUploadedFile(file, path)
	ctx.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// 将文件名映射到singer数据库中
	ctx.JSON(200,model.UpdateUserPic(userId,path[1:]))
}