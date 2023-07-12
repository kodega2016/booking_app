package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/kodega2016/booking-app/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("unable to get the requested template")
	}

	buff := new(bytes.Buffer)
	err := t.Execute(buff, nil)
	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buff.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}
}

/*
NOTE::simple way to create template cache

// tc holds the template cache
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if the template is already present in the cache,
	_, isOk := tc[t]

	if isOk {
		fmt.Println("template is already present in the cache,using from the cache")
	} else {
		// need to create template cache
		createTemplateCache(t)
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	tc[t] = tmpl
	return err
}


*/

func CreateTemplateCache() (map[string]*template.Template, error) {
	tc := map[string]*template.Template{}

	// get all the files name *.page.tmpl from "./templates"
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return tc, err
	}

	// range through all the files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return tc, err
		}

		// get all the layout files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return tc, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tc, err
			}
		}
		tc[name] = ts
	}

	return tc, nil
}
