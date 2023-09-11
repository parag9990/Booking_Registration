package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"BookingsWebsite/pkg/config"
	"BookingsWebsite/pkg/handlers"
	"BookingsWebsite/pkg/render"

	"github.com/alexedwards/scs"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// If we are in the development mode then we have to set the session false
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("Error creating template cache = ", err)
		log.Fatal(err)
	}
	app.TemplateCache = tc
	// If we are in the development mode then we need to load the template page instead of using the cached template
	app.UseCache = app.InProduction
	// It creates the New Repository Struct Instance
	repo := handlers.NewRepo(&app)
	// Pass the Repo Which is an instance of Repository to the New Handler in the Handler Package
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting Application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error While Starting Server = ", err)
	}
}
