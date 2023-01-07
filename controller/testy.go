package controller

import (
	"berpar/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"berpar/model"
)

type User struct {
	Name string `json:"username"`
	Password string `json:"password"`
}

type ResponseBody struct{
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func DefaultResponse() ResponseBody{
	user:=new(User)
	user.Name="Toma Chen"
	user.Password="chenxinyu"
	code:=294
	success:=true
	message:="Hello BerBer"
	return ResponseBody{
		Code: code,
		Success: success,
		Message: message,
		Data: user,
	}
}

func TestStringShowPass(ctx *gin.Context) {
	ctx.String(200, config.Configs.Database.DbPassword)
}

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,User{
		Name: ctx.Query("username"),
		Password: ctx.Query("password"),
	})
}

func TestGORM(ctx *gin.Context) {
	// res:=model.GetCategory(false)
	id,_:=strconv.ParseInt(ctx.Query("id"),10,64)
	res2:=model.SelectBySongId(uint(id))
	// ctx.JSON(200,res)
	ctx.JSON(200,res2)

	// ctx.String(200,res[0].Name)
}

func TestMaohao(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,User{
		Name: ctx.Param("username"),
		Password: ctx.Param("password"),
	})
}

func TestDefaultResponse(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,DefaultResponse())
}

