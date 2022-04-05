package handler

import (
	"net/http"
	"strconv"

	"github.com/erikrios/blog/handler/dummy"
	"github.com/erikrios/blog/handler/response"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	limitStr := queries.Get("limit")
	pageStr := queries.Get("page")

	var limit, page uint = 8, 1

	if val, err := strconv.Atoi(limitStr); err == nil {
		limit = uint(val)
	}

	if val, err := strconv.Atoi(pageStr); err == nil {
		page = uint(val)
	}

	articleList := dummy.GenerateDummyArticles(limit, page)

	response := response.Response[response.ArticleList]{
		Status:  "success",
		Message: "successfully get articles",
		Data:    articleList,
	}

	if err := tmpl.ExecuteTemplate(w, "index", response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong."))
		panic(err)
	}
}

func (h *HomeHandler) GetRobotsTxt(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User-agent: *\nAllow: /"))
}
