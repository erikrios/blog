package handler

import (
	"html/template"
	"os"
)

var tmpl *template.Template

func init() {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	htmlDir := rootDir + "/assets/html/"

	tmpl = template.Must(template.ParseGlob(htmlDir + "*.html"))
}
