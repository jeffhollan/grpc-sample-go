package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/jeffhollan/grpc-sample-go/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/xds"
)

func main() {
	addr, ok := os.LookupEnv("GRPC_SERVER_ADDRESS")
	opts := []grpc.DialOption{}

	if !ok {
		addr = "localhost:50051"
	}
	creds, err := xds.NewClientCredentials(xds.ClientOptions{
		FallbackCreds: insecure.NewCredentials(),
	})
	if err != nil {
		log.Fatalf("error creating client credentials: %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(addr, opts...)
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
