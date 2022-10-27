package service

import (
	"book_v3_go/dao"
	"book_v3_go/param"
	"book_v3_go/setting"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Index(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "欢迎使用！",
		})
	})
}

func Search(r *gin.Engine) {
	r.POST("/api/book/search", func(c *gin.Context) {
		var s param.SearchString
		c.ShouldBind(&s)
		limit, _ := strconv.Atoi(setting.Conf.WebConfig.SearchNum)
		c.JSON(200, gin.H{
			"data": dao.Search(s.Name, limit),
		})
	})
}

func Book(r *gin.Engine) {
	r.GET("/api/bookinfo/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, dao.Book(id))
	})
}

func Chapter(r *gin.Engine) {
	r.GET("/api/content/:id/:cid", func(c *gin.Context) {
		id := c.Param("id")
		cid := c.Param("cid")
		list := dao.Chapter(id, cid)
		if list == "" {
			c.JSON(500, "no")
		} else {
			c.JSON(200, list)
		}

	})
}

func Login(r *gin.Engine) {
	r.POST("/api/login", func(c *gin.Context) {
		var s param.UserString
		c.ShouldBind(&s)
		c.JSON(200, dao.Login(s.Id, s.Passwd))
	})
}

func BookCase(r *gin.Engine) {
	r.POST("/api/querybookshelf", func(c *gin.Context) {
		var s param.BookString
		c.ShouldBind(&s)
		c.JSON(200, dao.BookCase(s.Id))
	})
}

func UpBook(r *gin.Engine) {
	r.POST("/api/updatebookshelf", func(c *gin.Context) {
		var s param.UpBook
		c.ShouldBindJSON(&s)
		c.JSON(200, s.Bookshelf)
	})
}

func IndexHot(r *gin.Engine) {
	r.GET("/api/indexhot", func(c *gin.Context) {
		c.JSON(200, dao.IndexHot())
	})
}
func Rrootredirect(r *gin.Engine) {
	r.GET("/api/rootredirect", func(c *gin.Context) {
		c.String(200, "/book")
	})
}
func NoFound(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(500, "No")
	})
}
