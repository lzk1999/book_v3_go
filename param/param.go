package param

type SearchString struct {
	Name string `json:"search_string" form:"search_string"`
}

type UserString struct {
	Id     string `json:"id" form:"id"`
	Passwd string `json:"passwd" form:"passwd"`
}

type BookString struct {
	Id string `json:"id" form:"id"`
}

type UpBook struct {
	Bookshelf string `json:"bookshelf" form:"bookshelf"`
	Id        string `json:"id" form:"id"`
}
