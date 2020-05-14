package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) routes() *pat.PatternServeMux {
	mux := pat.New()

	// routes
	mux.Get("/", http.HandlerFunc(app.Index))

	// static files
	fileServer := http.FileServer(http.Dir(app.StaticDir))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
