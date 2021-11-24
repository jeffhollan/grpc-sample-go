package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	pb "github.com/jeffhollan/grpc-sample-go/protos"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
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
