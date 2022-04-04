package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/erikrios/blog/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	home := handler.NewHomeHandler()
	project := handler.NewProjectHandler()
	resume := handler.NewResumeHandler()
	about := handler.NewAboutHandler()

	r.Get("/", home.GetHome)
	r.Get("/project", project.GetProject)
	r.Get("/resume", resume.GetResume)
	r.Get("/about", about.GetAbout)
	r.Get("/robots.txt", home.GetRobotsTxt)

	fs := http.FileServer(http.Dir("./assets/statics"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	log.Printf("Server listening on port %s...\n", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
