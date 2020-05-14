package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

// PrepareTemplate ...takes in name of template and the layout it inherits from, returning the template.
func (app *App) PrepareTemplate(fileName, layout string) (*template.Template, error) {
	// Default funcs available to views that are prepared using Prepared Template.
	defaultViewFuncs := template.FuncMap{
		// Usage... <img src="{{ ImgSrc "imageName" }}">
		"ImgSrc": func(imgName string) string {
			return fmt.Sprintf("static/images/%s", imgName)
		},

		// Usage... <script src="{{ ReactComponent "nameOfComponent" }}"></script>
		"ReactComponent": func(componentName string) string {
			return fmt.Sprintf("static/javascripts/dist/components/%s/%s.js", componentName, componentName)
		},
	}

	templatePath := fmt.Sprintf("%s/%s.gohtml", app.TemplatesDir, fileName)
	templateLayout := fmt.Sprintf("%s/layouts/%s.gohtml", app.TemplatesDir, layout)
	templateName := strings.TrimPrefix(templatePath, fmt.Sprintf("%s/", app.TemplatesDir))
	template, err := template.New(templateName).Funcs(defaultViewFuncs).ParseFiles(templatePath, templateLayout)
	if err != nil {
		return nil, err
	}

	return template, nil
}

// ServerError ...writes stack trace and returns error response
func (app *App) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// ViewData ...returns data to be used in view; comes with default data.
func (app *App) ViewData(view string, data map[string]interface{}) map[string]interface{} {
	if data == nil {
		data = make(map[string]interface{})
	}

	convertToProjectPath := func(url string) string {
		return strings.Replace(url, "static", app.StaticDir, 1)
	}

	defaultData := make(map[string]interface{})
	pathToJavaScript := fmt.Sprintf("static/javascripts/%s/%s.js", view, view)
	pathToMainJS := "static/javascripts/dist/main.js"
	pathToReact := "static/javascripts/node_modules/react/umd/react.development.js"
	pathToReactDOM := "static/javascripts/node_modules/react-dom/umd/react-dom.development.js"
	pathToStylesheet := fmt.Sprintf("static/css/%s.css", view)

	defaultDataPaths := map[string]string{
		"CSS":        pathToStylesheet,
		"JavaScript": pathToJavaScript,
		"React":      pathToReact,
		"ReactDOM":   pathToReactDOM,
		"MainJS":     pathToMainJS,
	}

	for k, v := range defaultDataPaths {
		if _, err := os.Stat(convertToProjectPath(v)); !os.IsNotExist(err) {
			defaultData[k] = v
		}
	}

	if len(defaultData) > 0 {
		for k, v := range defaultData {
			data[k] = v
		}
	}

	return data
}
