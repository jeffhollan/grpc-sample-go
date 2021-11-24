package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	pb "github.com/jeffhollan/grpc-sample-go/protos"
)

// App export
type App struct {
	Router *mux.Router
	Client pb.GreeterClient
}

// TODO: add handlers here
func (app *App) helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	grpc_r, err := app.Client.SayHello(ctx, &pb.HelloRequest{Name: "Azure Container Apps"})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	w.Write([]byte(grpc_r.GetMessage()))
}

func (app *App) initializeRoutes(Client pb.GreeterClient) {
	app.Client = Client
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", app.helloHandler)
}

func (app *App) run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
