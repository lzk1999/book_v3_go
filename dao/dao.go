package dao

import (
	"book_v3_go/dto"
	"book_v3_go/model"
	"book_v3_go/mysql"
	"book_v3_go/param"
	"book_v3_go/utils"
)

func Search(name string, limit int) (list interface{}) {
	var alist []model.Article
	mysql.DB.Model(&model.Article{}).Where("articlename LIKE ?", "%"+name+"%").Limit(limit).Scan(&alist)
	articles := dto.ArticleDtos(alist)
	if articles != nil {
		return articles
	}
	return ""
}

func Book(id string) (list interface{}) {
	var article model.Article
	var alist []model.ArticleChapter
	mysql.DB.First(&article, id)
	mysql.DB.Table(utils.TableName(id)).Model(&model.ArticleChapter{}).Where("articleid = ?", id).Order("chapterorder asc").Scan(&alist)
	href, text := dto.BookChapterDto(alist)
	return dto.BookArticleDto(article, href, text)
}

func Chapter(id string, cid string) (list interface{}) {
	var alist model.ArticleChapter
	mysql.DB.Table(utils.TableName(id)).Model(&model.ArticleChapter{}).Where("articleid = ? AND chapterid = ?", id, cid).Scan(&alist)
	return dto.ChapterDto(alist)
}

func Login(id string, passwd string) (list interface{}) {
	var alist *model.SystemUsers
	var clist *model.SystemUsers
	mysql.DB.Model(&model.SystemUsers{}).Where("uname = ? ", id).Scan(&alist)
	if alist == nil {
		return model.Login{
			Context: "当前用户未注册。",
			Res:     false,
		}
	}
	mysql.DB.Model(&model.SystemUsers{}).Where("uname = ? AND pass = ? ", id, utils.PassSalt(passwd, alist.Salt)).Scan(&clist)
	if clist != nil {
		return dto.LoginDto(*clist)
	} else {
		return model.Login{
			Context: "输入密码有误。",
			Res:     false,
		}
	}
}

func BookCase(id string) (list interface{}) {
	var alist []model.ArticleBookcase
	mysql.DB.Model(&model.ArticleBookcase{}).Where("username = ?", id).Scan(&alist)
	if alist != nil {
		return dto.BookCase(alist)
	}
	return model.BookCase{
		Bookshelf: "{}",
		Res:       true,
	}
}

func IndexHot() (list interface{}) {
	var alist []model.Article
	mysql.DB.Model(&model.Article{}).Limit(6).Scan(&alist)
	articles := dto.ArticleDtos(alist)
	if articles != nil {
		return articles
	}
	return ""
}

func UpBook(bookshelf param.UpBook, id string) (list interface{}) {
	var user model.SystemUsers
	mysql.DB.Model(&model.SystemUsers{}).Where("uname = ?", id).Scan(&user)
	//for _, m := range alist {
	//	c := model.ArticleBookcase{
	//		Articleid:   m.Num2.BookID,
	//		Articlename: m.Num2.Name,
	//		Userid:      user.Uid,
	//		Username:    user.Uname,
	//		Chapterid:   m.Num2.CatalogueID,
	//		Chaptername: m.Num2.CatalogueName,
	//	}
	//	mysql.DB.Create(&c)
	//}
	return model.ChapterBookshelf{
		Context: "更新书架成功。",
		Res:     true,
	}
}
