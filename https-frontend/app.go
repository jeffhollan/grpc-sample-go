package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App export
type App struct {
	Router *mux.Router
}

// TODO: add handlers here
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Mux!"))
}

func (app *App) initializeRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", helloHandler)
}

func (app *App) run() {
	log.Fatal(http.ListenAndServe(":8050", app.Router))
}
