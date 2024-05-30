package main

import (
	"html/template"
	"net/http"
	"path"
)

var pathToTemplates = "./templates/"

func (app *application) Home(writer http.ResponseWriter, request *http.Request) {
	//fmt.Fprint(writer, "This is the home page.")
	_ = app.render(writer, request, "home.page.gohtml", &TemplateData{})
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(writer http.ResponseWriter, request *http.Request, templt string, data *TemplateData) error {
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, templt))
	if err != nil {
		http.Error(writer, "bad request", http.StatusBadRequest)
		return err
	}

	data.IP = app.ipFromContext(request.Context())

	err = parsedTemplate.Execute(writer, data)

	if err != nil {
		return err
	}
	return nil
}
