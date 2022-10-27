package dto

import (
	"book_v3_go/model"
	"book_v3_go/mysql"
	"book_v3_go/utils"
	"fmt"
	"github.com/goccy/go-json"
)

func BookCase(alist []model.ArticleBookcase) (list model.BookCase) {
	var blist model.Article
	var bshelf []model.Bookshelf
	for _, bookcase := range alist {
		mysql.DB.Model(&model.Article{}).Where("articleid = ?", bookcase.Articleid).Scan(&blist)
		caser := BookCaser(bookcase, blist)
		bshelf = append(bshelf, caser)
	}
	jsons, _ := json.Marshal(bshelf) //转换成JSON返回的是byte[]
	return model.BookCase{
		Bookshelf: string(jsons),
		Res:       true,
	}

}

func BookCaser(alist model.ArticleBookcase, blist model.Article) (list model.Bookshelf) {
	return model.Bookshelf{
		Id:            alist.Caseid,
		BookId:        alist.Articleid,
		Name:          alist.Articlename,
		ImgUrl:        utils.Img(blist.Rgroup, alist.Articleid),
		Author:        blist.Author,
		CatalogueId:   alist.Chapterid,
		CatalogueName: alist.Chaptername,
	}
}

func LoginDto(alist model.SystemUsers) (list model.Login) {
	return model.Login{
		Context:  "登录成功。",
		FaceId:   false,
		Res:      true,
		UserName: alist.Name,
	}
}

func ChapterDto(alist model.ArticleChapter) model.BookChapter {
	result, err := utils.Txt(alist.Articleid, alist.Chapterid)
	fmt.Println()
	if err != nil {
		result = "获取文本错误，请检查配置"
	}
	return model.BookChapter{
		CatalogueName: alist.Chaptername,
		Content:       result,
	}

}

func BookChapterDto(chpaters []model.ArticleChapter) (href []string, text []string) {
	for _, chpater := range chpaters {
		href = append(href, fmt.Sprintf("/book/%d/%d", chpater.Articleid, chpater.Chapterid))
		text = append(text, chpater.Chaptername)
	}
	return href, text
}

func BookArticleDto(article model.Article, href []string, text []string) model.BookArticle {
	return model.BookArticle{
		Author:            article.Author,
		BookId:            article.Articleid,
		ImgUrl:            utils.Img(article.Rgroup, article.Articleid),
		Name:              article.Articlename,
		Resume:            article.Intro,
		CatalogueHrefList: href,
		CatalogueTextList: text,
	}
}

func ArticleDto(article model.Article) model.ArticleFormated {
	return model.ArticleFormated{
		Author: article.Author,
		BookId: article.Articleid,
		ImgUrl: utils.Img(article.Rgroup, article.Articleid),
		Name:   article.Articlename,
		Resume: article.Intro,
	}
}
func ArticleDtos(article []model.Article) (articles []model.ArticleFormated) {
	for _, a := range article {
		c := ArticleDto(a)
		articles = append(articles, c)
	}
	return articles

}
