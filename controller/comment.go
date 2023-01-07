package controller

import (
	"berpar/model"
	"strconv"

	"github.com/gin-gonic/gin"
)


// @Summary  提交评论
// @Produce json
// @Tags  comment
// @Router /api/comment/add [post]
func AddComment(ctx *gin.Context){
	userId,_:=strconv.Atoi(ctx.PostForm("userId"))
	songListId,_:=strconv.Atoi(ctx.PostForm("songListId"))
	songId,_:=strconv.Atoi(ctx.PostForm("songId"))
	content:=ctx.PostForm("content")
	comment:=model.Comment{UserId: uint(userId),SongId: uint(songId),
	SongListId: uint(songListId),Content: content}
	ctx.JSON(200,model.AddComment(comment))
}

// @Summary  删除评论
// @Produce json
// @Tags  comment
// @Router /api/comment/delete [get]
func DeleteComment(ctx *gin.Context){
	id,_:=strconv.Atoi(ctx.Query("id"))
	ctx.JSON(200,model.DeleteComment(uint(id)))
}

// @Summary  获得指定歌曲ID的评论列表
// @Produce json
// @Tags  comment
// @Router /api/comment/song/detail [get]
func CommentOfSongId(ctx *gin.Context){
	songId,_:=strconv.Atoi(ctx.Query("songId"))
	ctx.JSON(200,model.CommentOfSongId(uint(songId)))
}

// @Summary  获得指定歌单ID的评论列表
// @Produce json
// @Tags  comment
// @Router /api/comment/songList/detail [get]
func CommentOfSongListId(ctx *gin.Context){
	songListId,_:=strconv.Atoi(ctx.Query("songListId"))
	ctx.JSON(200,model.CommentOfSongListId(uint(songListId)))
}

// @Summary  点赞
// @Produce json
// @Tags  comment
// @Router /api/comment/like [post]
func UpdateCommentMsg(ctx *gin.Context){
	id,_:=strconv.Atoi(ctx.PostForm("id"))
	up,_:=strconv.Atoi(ctx.PostForm("up"))
	ctx.JSON(200,model.UpdateCommentMsg(id,up))
}