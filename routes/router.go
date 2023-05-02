package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"


	"berpar/config"
	"berpar/controller"

	_ "berpar/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	r := gin.Default()

	// 中间件
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Use(cors.Default())
	store := cookie.NewStore([]byte("secret11111"))
	r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("api")

	// 静态资源
	v1.Static("/img", "./img")
	v1.Static("/songs", "./song") //这里由原先的song变成了songs

	v1.GET("/", controller.TestStringShowPass)
	v1.GET("/hello", controller.Hello)
	v1.GET("/v1", controller.TestGORM)
	v1.GET("/hello/:username/:password", controller.TestMaohao)
	v1.GET("/dr", controller.TestDefaultResponse)

	admin := v1.Group("/admin")
	{
		admin.POST("/login/status", controller.LoginStatus)
	}

	collect := v1.Group("/collection")
	{
		collect.POST("/add", controller.AddCollection)
		collect.POST("/delete", controller.DeleteCollect)
		collect.POST("/status", controller.IsCollection)
		collect.GET("/detail", controller.CollectionOfUser)
	}

	comment := v1.Group("/comment")
	{
		comment.POST("/add", controller.AddComment)
		comment.GET("/delete", controller.DeleteComment)
		comment.GET("/song/detail", controller.CommentOfSongId)
		comment.GET("/songList/detail", controller.CommentOfSongListId)
		comment.POST("/like", controller.UpdateCommentMsg)
	}

	user := v1.Group("/user")
	{
		user.POST("/add", controller.AddUser)
		user.POST("/login/status", controller.LoginStatus)
		user.GET("", controller.GetAllUser)
		user.GET("/detail",controller.SelectByUserId)
		user.GET("/delete", controller.DeleteUser)
		user.POST("/update", controller.UpdateUserMsg)
		user.POST("/updatePassword", controller.UpdateUserPassword)
		user.POST("/avator/update", controller.UpdateUserPic)
	}

	listSong := v1.Group("/listSong")
	{
		listSong.POST("/add", controller.AddListSong)
		listSong.GET("/delete", controller.DeleteListSong)
		listSong.GET("/detail", controller.ListSongOfId)
		listSong.POST("/update", controller.UpdateListSongMsg)
	}

	rankList := v1.Group("/rankList")
	{
		rankList.GET("", controller.GetScoreBySongList)
		rankList.GET("/user", controller.GetScoreByUser)
		rankList.POST("/add", controller.AddScore)
	}

	singer := v1.Group("/singer")
	{
		singer.GET("", controller.GetSinger)
		singer.POST("/update", controller.UpdateSingerMsg)
		singer.POST("/add", controller.AddSinger)
		singer.GET("/delete", controller.DeleteSinger)
		singer.GET("/name/detail", controller.SingerOfName)
		singer.GET("/sex/detail", controller.SingerOfSex)
		singer.POST("/avatar/update", controller.UpdateSingerPic)
	}

	song := v1.Group("/song")
	{
		song.GET("/detail", controller.SelectBySongId)
		song.GET("/singer/detail", controller.SelectBySingerId)
		song.GET("/singerName/detail", controller.SelectBySingerName)
		song.GET("", controller.GetAllSong)
		song.POST("/add", controller.AddSong)
		song.GET("/delete", controller.DeleteSong)
		song.POST("/update", controller.UpdateSongMsg)
		song.POST("/img/update", controller.UpdateSongPic)
		song.POST("/url/update", controller.UpdateSongUrl)
		song.GET("/modify",controller.ModifySongUrl)
	}

	songlist := v1.Group("/songList")
	{
		songlist.POST("/add", controller.AddSongList)
		songlist.GET("/delete", controller.DeleteSongList)
		songlist.GET("", controller.GetAllSongList)
		songlist.POST("/update", controller.UpdateSongListMsg)
		songlist.POST("/img/update", controller.UpdateSongListPic)
		songlist.GET("/style/detail", controller.SongListOfStyle)
		songlist.GET("/likeTitle/detail", controller.SongListOfLikeTitle)
	}

	r.Run(config.Configs.Server.Port)
}


func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        origin := c.Request.Header.Get("Origin") //请求头部
        if origin != "" {
            //接收客户端发送的origin （重要！）
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin) 
            //服务器支持的所有跨域请求的方法
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") 
            //允许跨域设置可以返回其他子段，可以自定义字段
            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X\\_Requested\\_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
            // 允许浏览器（客户端）可以解析的头部 （重要）
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers") 
            //设置缓存时间
            c.Header("Access-Control-Max-Age", "172800") 
            //允许客户端传递校验信息比如 cookie (重要)
            c.Header("Access-Control-Allow-Credentials", "true")                                                                            
        }

        //允许类型校验 
        if method == "OPTIONS" {
            c.JSON(http.StatusOK, "ok!")
        }

        // defer func() {
        //     if err := recover(); err != nil {
        //         log.Printf("Panic info is: %v", err)
        //     }
        // }()
        c.Next()
    }
}

func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}