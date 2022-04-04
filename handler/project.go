package handler

import (
	"net/http"
)

type ProjectHandler struct{}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}

func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "project", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong."))
		panic(err)
	}
}
