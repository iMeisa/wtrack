package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/iMeisa/errortrace"
	"github.com/iMeisa/weed/internal/config"
	"github.com/iMeisa/weed/internal/handlers"
	"github.com/iMeisa/weed/internal/models"
	"github.com/iMeisa/weed/internal/render"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Load env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// WebApp settings
	app.Prod = os.Getenv("ENV") == "prod"

	session = scs.New()
	session.Lifetime = 24 * time.Hour * 90 // 90 days
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Prod

	app.Session = session

	// Create json db
	if _, err := os.Stat("./db/db.json"); os.IsNotExist(err) {
		db := models.DB{}

		// Marshal the leagues into JSON
		dbJSON, err := json.MarshalIndent(db, "", "	")
		if err != nil {
			trace := errortrace.NewTrace(err)
			trace.Read()
			panic(err)
		}

		// Write the JSON to the file
		err = ioutil.WriteFile(fmt.Sprintf("./db/db.json"), dbJSON, 0644)
		if err != nil {
			trace := errortrace.NewTrace(err)
			trace.Read()
			panic(err)
		}
	}

	// Templates
	tc, trace := render.CreateTemplateCache()
	if trace.HasError() {
		trace.Read()
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)

	// Serve
	fmt.Println(fmt.Sprintf("Starting %s application on port %s", os.Getenv("ENV"), os.Getenv("SITE_PORT")))

	srv := &http.Server{
		Addr:    os.Getenv("SITE_PORT"),
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		trace = errortrace.NewTrace(err)
		trace.Read()
		log.Fatal(err)
	}
}
