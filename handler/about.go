package handler

import "net/http"

type AboutHandler struct{}

func NewAboutHandler() *AboutHandler {
	return &AboutHandler{}
}

func (a *AboutHandler) GetAbout(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "about", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong."))
		panic(err)
	}
}
