package main

import (
	"net/http"
)

//Index ...homepage.
func (app *App) Index(w http.ResponseWriter, r *http.Request) {
	templates, err := app.PrepareTemplate("index", "application")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = templates.Execute(w, app.ViewData("index", nil))
	if err != nil {
		app.ServerError(w, err)
		return
	}
}
