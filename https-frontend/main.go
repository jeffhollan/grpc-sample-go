package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/jeffhollan/grpc-sample-go/protos"
	"google.golang.org/grpc"
)

func main() {
	addr, ok := os.LookupEnv("GRPC_SERVER_ADDRESS")

	if !ok {
		addr = "localhost:50051"
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	app := App{}
	app.initializeRoutes(
		c,
	)
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8050"
	}
	binding := fmt.Sprintf(":%s", port)

	app.run(binding)
}
