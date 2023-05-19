package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	smtp struct {
		host     string
		port     int
		username string
		password string
	}
	frontend string
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 5000, "Server port to listen on")
	flag.StringVar(&cfg.smtp.host, "smtphost", "sandbox.smtp.mailtrap.io", "smtp host")
	flag.IntVar(&cfg.smtp.port, "smtpport", 587, "smtp port")
	flag.StringVar(&cfg.smtp.username, "smtpuser", os.Getenv("smtpuser"), "smtp username")
	flag.StringVar(&cfg.smtp.password, "smtppassword", os.Getenv("smtppassword"), "smtp password")
	flag.StringVar(&cfg.frontend, "frontend", "http://localhost:4000", "url to frontend")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
	}

	app.CreateDirIfNotExist("./invoices")

	err := app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Println(fmt.Sprintf("Starting invoice microservice on port %d", app.config.port))

	return srv.ListenAndServe()
}
