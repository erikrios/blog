package handler

import (
	"net/http"
)

type ResumeHandler struct{}

func NewResumeHandler() *ResumeHandler {
	return &ResumeHandler{}
}

func (h *ResumeHandler) GetResume(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "resume", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong."))
		panic(err)
	}
}
