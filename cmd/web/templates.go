package main

import ( 
	"fmt"
	"snippetbox/internal/models"
	"path/filepath"
	"html/template"

)


type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}


func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	fmt.Println(pages)

	for _, page := range pages {
		name := filepath.Base(page)
		
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		fmt.Println("ts:", ts)

		cache[name] = ts
	}
	
	fmt.Println("cache:", cache)	
	return cache, nil
	
}









