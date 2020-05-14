package main

import (
	"log"
	"net/http"
	"os"
)

// App ...struct for app dependencies
type App struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	StaticDir    string
	TemplatesDir string
}

var (
	app    App
	server *http.Server
)

func init() {
	app = App{
		InfoLog:      log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
		ErrorLog:     log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
		StaticDir:    "./ui/static",
		TemplatesDir: "./ui/templates",
	}

	server = &http.Server{
		Addr:     ":8080",
		ErrorLog: app.ErrorLog,
		Handler:  app.routes(),
	}
}
