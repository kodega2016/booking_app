package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	if err != nil {
		log.Fatal(err)
	}
	parsedTemplate.Execute(w, nil)
}
