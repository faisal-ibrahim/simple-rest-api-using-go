package main

import (
	"fmt"
	"github.com/github.com/faisal-ibrahim/simple-rest-api-using-go/internal/comment"
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
