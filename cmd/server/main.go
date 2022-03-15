package main

import "fmt"

// App - the struct which contains things like pointers to the database connections
type App struct {
}

// Run - handles the startup of our application
func (app *App) Run() error {
	fmt.Println("Setting up our App!")
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
