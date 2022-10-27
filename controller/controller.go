package controller

import (
	"book_v3_go/service"
	"github.com/gin-gonic/gin"
)

func Index(r *gin.Engine) {
	service.Index(r)
}

func Search(r *gin.Engine) {
	service.Search(r)
}

func Book(r *gin.Engine) {
	service.Book(r)
}

func Chapter(r *gin.Engine) {
	service.Chapter(r)
}

func Login(r *gin.Engine) {
	service.Login(r)
}

func BookCase(r *gin.Engine) {
	service.BookCase(r)
}

func UpBook(r *gin.Engine) {
	service.UpBook(r)
}

func IndexHot(r *gin.Engine) {
	service.IndexHot(r)
}
func NoFound(r *gin.Engine) {
	service.NoFound(r)
}

func Rrootredirect(r *gin.Engine) {
	service.Rrootredirect(r)
}
