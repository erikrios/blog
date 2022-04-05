package dummy

import (
	"fmt"

	"github.com/erikrios/blog/handler/response"
)

func GenerateDummyArticles(limit uint, page uint) response.ArticleList {
	articles := make([]response.Article, 0)

	for i := 0; uint(i) < limit; i++ {
		article := response.Article{
			ID:            fmt.Sprintf("articles-%d", i+1),
			Title:         fmt.Sprintf("This is Article Title %d", i+1),
			PublishedDate: fmt.Sprintf("%d %s %d", i+1, "April", 2022),
			Author:        fmt.Sprintf("Erik %d", i+1),
			Content:       fmt.Sprintf("%d %s", i+1, "Lorem, ipsum dolor sit amet consectetur adipisicing elit.\nDistinctio, veritatis voluptatum molestiae numquam ducimus eligendi sunt mollitia porro nisi voluptas. Cum odio ratione, reprehenderit totam dolorum libero itaque cumque quo? Lorem ipsum dolor sit amet consectetur adipisicing elit. Doloribus repellat vero provident dignissimos reiciendis odio tenetur consequuntur soluta, blanditiis, et similique nihil velit. Ipsum doloremque atque dolores officiis optio. Mollitia."),
		}
		articles = append(articles, article)
	}

	pageInfo := response.PageInfo{
		Limit:     limit,
		Page:      page,
		TotalPage: 3,
		Total:     limit * 3,
	}

	pageInfo.SetPrevPage()
	pageInfo.SetNextPage()

	articleList := response.ArticleList{
		Articles: articles,
		PageInfo: pageInfo,
	}

	return articleList
}
