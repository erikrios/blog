package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erikrios/blog/config"
	"github.com/erikrios/blog/constant"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	port := constant.PORT
	log.Printf("Server starting at port :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
