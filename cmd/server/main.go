package main

import (
	"fmt"
	"github.com/github.com/faisal-ibrahim/simple-rest-api-using-go/internal/database"
	transportHTTP "github.com/github.com/faisal-ibrahim/simple-rest-api-using-go/internal/transport/http"
	"net/http"
)

// App - the struct which contains things like pointers to the database connections
type App struct {
}

// Run - handles the startup of our application
func (app *App) Run() error {
	fmt.Println("Setting up our App!")
	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}
func main() {
	fmt.Println("Simple REST API using GO!")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting our REST API")
		fmt.Println(err)
	}
}
