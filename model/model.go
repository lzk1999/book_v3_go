package model

type ChapterBookshelf struct {
	Context string `json:"context"`
	Res     bool   `json:"res"`
}

type Bookshelf1 struct {
	Bookshelf1 struct {
		BookId        int64  `json:"book_id" form:"book_id"`
		Name          string `json:"name" form:"name"`
		ImgUrl        string `json:"img_url" form:"img_url"`
		Author        string `json:"author" form:"author"`
		CatalogueId   int64  `json:"catalogue_id" form:"catalogue_id"`
		CatalogueName string `json:"catalogue_name" form:"catalogue_name"`
	} `json:"" form:""`
}

type Bookshelf struct {
	Id            int64  `json:"id"`
	BookId        int64  `json:"book_id"`
	Name          string `json:"name"`
	ImgUrl        string `json:"img_url"`
	Author        string `json:"author"`
	CatalogueId   int64  `json:"catalogue_id"`
	CatalogueName string `json:"catalogue_name"`
}

type BookCase struct {
	Bookshelf string `json:"bookshelf"`
	Res       bool   `json:"res"`
}

type ArticleBookcase struct {
	Caseid      int64  `gorm:"column:caseid" db:"caseid"`
	Articleid   int64  `gorm:"column:articleid" db:"articleid"`
	Articlename string `gorm:"column:articlename" db:"articlename"`
	Userid      int64  `gorm:"column:userid" db:"userid"`
	Username    string `gorm:"column:username" db:"username"`
	Chapterid   int64  `gorm:"column:chapterid" db:"chapterid"`
	Chaptername string `gorm:"column:chaptername" db:"chaptername"`
}

func (ArticleBookcase) TableName() string {
	return "shipsay_article_bookcase"
}

type Login struct {
	Context  string `json:"context"`
	FaceId   bool   `json:"faceid"`
	Res      bool   `json:"res"`
	UserName string `json:"user_name"`
}

type SystemUsers struct {
	Uid       int64  `gorm:"column:uid" db:"uid"`
	Uname     string `gorm:"column:uname" db:"uname"`
	Name      string `gorm:"column:name" db:"name"`
	Pass      string `gorm:"column:pass" db:"pass"`
	Salt      string `gorm:"column:salt" db:"salt"`
	Groupid   int64  `gorm:"column:groupid" db:"groupid"`
	Regdate   int64  `gorm:"column:regdate" db:"regdate"`
	Lastlogin int64  `gorm:"column:lastlogin" db:"lastlogin"`
	Email     string `gorm:"column:email" db:"email"`
}

func (SystemUsers) TableName() string {
	return "shipsay_system_users"
}

type BookChapter struct {
	CatalogueName string `json:"catalogue_name"`
	Content       string `json:"content"`
}

type BookArticle struct {
	Author            string   `json:"book_author"`
	BookId            int64    `json:"book_id"`
	ImgUrl            string   `json:"book_img_url"`
	Name              string   `json:"book_name"`
	Resume            string   `json:"book_resume"`
	CatalogueHrefList []string `json:"catalogue_href_list"`
	CatalogueTextList []string `json:"catalogue_text_list"`
}

type ArticleChapter struct {
	Chapterid    int64  `gorm:"column:chapterid" db:"chapterid"`
	Articleid    int64  `gorm:"column:articleid" db:"articleid"`
	Articlename  string `gorm:"column:articlename" db:"articlename"`
	Posterid     int64  `gorm:"column:posterid" db:"posterid"`
	Poster       string `gorm:"column:poster" db:"poster"`
	Postdate     int64  `gorm:"column:postdate" db:"postdate"`
	Lastupdate   int64  `gorm:"column:lastupdate" db:"lastupdate"`
	Chaptername  string `gorm:"column:chaptername" db:"chaptername"`
	Chapterorder int64  `gorm:"column:chapterorder" db:"chapterorder"`
	Words        int64  `gorm:"column:words" db:"words"`
	Chaptertype  int64  `gorm:"column:chaptertype" db:"chaptertype"`
	Attachment   string `gorm:"column:attachment" db:"attachment"`
	Summary      string `gorm:"column:summary" db:"summary"`
	Isimage      int64  `gorm:"column:isimage" db:"isimage"`
	Volumeid     int64  `gorm:"column:volumeid" db:"volumeid"`
	Pushed       int64  `gorm:"column:pushed" db:"pushed"`
}

type ArticleFormated struct {
	Author string `json:"author"`
	BookId int64  `json:"book_id"`
	ImgUrl string `json:"img_url"`
	Name   string `json:"name"`
	Resume string `json:"resume"`
}

type Article struct {
	Articleid     int64  `gorm:"column:articleid" db:"articleid"`
	Siteid        int64  `gorm:"column:siteid" db:"siteid"`
	Postdate      int64  `gorm:"column:postdate" db:"postdate"`
	Lastupdate    int64  `gorm:"column:lastupdate" db:"lastupdate"`
	Infoupdate    int64  `gorm:"column:infoupdate" db:"infoupdate"`
	Articlename   string `gorm:"column:articlename" db:"articlename"`
	Articlecode   string `gorm:"column:articlecode" db:"articlecode"`
	Backupname    string `gorm:"column:backupname" db:"backupname"`
	Keywords      string `gorm:"column:keywords" db:"keywords"`
	Roles         string `gorm:"column:roles" db:"roles"`
	Initial       string `gorm:"column:initial" db:"initial"`
	Authorid      int64  `gorm:"column:authorid" db:"authorid"`
	Author        string `gorm:"column:author" db:"author"`
	Posterid      int64  `gorm:"column:posterid" db:"posterid"`
	Poster        string `gorm:"column:poster" db:"poster"`
	Sortid        int64  `gorm:"column:sortid" db:"sortid"`
	Typeid        int64  `gorm:"column:typeid" db:"typeid"`
	Intro         string `gorm:"column:intro" db:"intro"`
	Lastchapterid int64  `gorm:"column:lastchapterid" db:"lastchapterid"`
	Lastchapter   string `gorm:"column:lastchapter" db:"lastchapter"`
	Lastsummary   string `gorm:"column:lastsummary" db:"lastsummary"`
	Chapters      int64  `gorm:"column:chapters" db:"chapters"`
	Words         int64  `gorm:"column:words" db:"words"`
	Lastvisit     int64  `gorm:"column:lastvisit" db:"lastvisit"`
	Dayvisit      int64  `gorm:"column:dayvisit" db:"dayvisit"`
	Weekvisit     int64  `gorm:"column:weekvisit" db:"weekvisit"`
	Monthvisit    int64  `gorm:"column:monthvisit" db:"monthvisit"`
	Allvisit      int64  `gorm:"column:allvisit" db:"allvisit"`
	Lastvote      int64  `gorm:"column:lastvote" db:"lastvote"`
	Dayvote       int64  `gorm:"column:dayvote" db:"dayvote"`
	Weekvote      int64  `gorm:"column:weekvote" db:"weekvote"`
	Monthvote     int64  `gorm:"column:monthvote" db:"monthvote"`
	Allvote       int64  `gorm:"column:allvote" db:"allvote"`
	Goodnum       int64  `gorm:"column:goodnum" db:"goodnum"`
	Fullflag      int64  `gorm:"column:fullflag" db:"fullflag"`
	Imgflag       int64  `gorm:"column:imgflag" db:"imgflag"`
	Display       int64  `gorm:"column:display" db:"display"`
	Rgroup        int64  `gorm:"column:rgroup" db:"rgroup"`
	Lastvolume    string `gorm:"column:lastvolume" db:"lastvolume"`
	Lastvolumeid  int64  `gorm:"column:lastvolumeid" db:"lastvolumeid"`
	Pushed        int64  `gorm:"column:pushed" db:"pushed"`
	Ratenum       int64  `gorm:"column:ratenum" db:"ratenum"`
	Ratesum       int64  `gorm:"column:ratesum" db:"ratesum"`
}

func (Article) TableName() string {
	return "shipsay_article_article"
}
