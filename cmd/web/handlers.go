package main

import (
	"html/template"
	"net/http"
	"path"
)

var pathToTemplates = "./templates/"

func (app *application) Home(writer http.ResponseWriter, reader *http.Request) {
	//fmt.Fprint(writer, "This is the home page.")
	_ = app.render(writer, reader, "home.page.gohtml", &TemplateData{})
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(writer http.ResponseWriter, reader *http.Request, templt string, data *TemplateData) error {
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, templt))
	if err != nil {
		http.Error(writer, "bad request", http.StatusBadRequest)
		return err
	}

	err = parsedTemplate.Execute(writer, data)

	if err != nil {
		return err
	}
	return nil
}
