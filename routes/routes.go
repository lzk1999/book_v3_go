package routes

import (
	"book_v3_go/controller"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r = gin.New()
	//首页路由注册
	controller.Index(r)
	controller.Search(r)
	controller.Book(r)
	controller.Chapter(r)
	controller.Login(r)
	controller.BookCase(r)
	controller.UpBook(r)
	controller.IndexHot(r)
	controller.NoFound(r)
	controller.Rrootredirect(r)
	return r
}
