package main

import (
	"github.com/github.com/faisal-ibrahim/simple-rest-api-using-go/internal/comment"
	"github.com/github.com/faisal-ibrahim/simple-rest-api-using-go/internal/database"
	transportHTTP "github.com/github.com/faisal-ibrahim/simple-rest-api-using-go/internal/transport/http"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// App - application information
type App struct {
	Name    string
	Version string
}

// Run - handles the startup of our application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up application")
	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	return nil
}
func main() {
	app := App{
		Name:    "REST comment API service",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error("Error starting our REST API")
		log.Fatal(err)
	}
}
