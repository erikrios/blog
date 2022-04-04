package handler

import (
	"net/http"
)

type ContactHandler struct{}

func NewContactHandler() *ContactHandler {
	return &ContactHandler{}
}

func (h *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "contact", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong."))
		panic(err)
	}
}
