package router

//路由
import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	//e.GET("/index", controller.ListUser)
	e.POST("/register", controller.Register)
	e.GET("/register", controller.GoRegister)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.GET("/", controller.Index)

	//操作博客
	//博客列表
	e.GET("/post_index", controller.GetPostIndex)
	//添加博客
	e.POST("/post", controller.AddPost)
	//跳转到添加博客页面
	e.GET("/post", controller.GoAddPost)
	e.GET("/postdetail", controller.PostDetail)

	e.Run(":8888")
}
