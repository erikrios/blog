package response

type Response[T any] struct {
	Status  string
	Message string
	Data    T
}

type Article struct {
	ID            string
	Title         string
	PublishedDate string
	Author        string
	Content       string
	TotalViews    uint
}

type PageInfo struct {
	Limit     uint
	Page      uint
	TotalPage uint
	Total     uint
	NextPage  uint
	PrevPage  uint
}

func (p *PageInfo) SetNextPage() {
	if p.Page != p.TotalPage {
		p.NextPage = p.Page + 1
	} else {
		p.NextPage = p.TotalPage
	}
}

func (p *PageInfo) SetPrevPage() {
	if p.Page > 1 {
		p.PrevPage = p.Page - 1
	} else {
		p.PrevPage = 1
	}
}

type ArticleList struct {
	Articles        []Article
	PageInfo        PageInfo
	PopularArticles []Article
}
