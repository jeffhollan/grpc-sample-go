package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/jeffhollan/grpc-sample-go/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	addr, ok := os.LookupEnv("GRPC_SERVER_ADDRESS")
	opts := []grpc.DialOption{}
	if !ok {
		addr = "localhost:50051"
		opts = append(opts, grpc.WithInsecure())
	} else {
		config := &tls.Config{
			InsecureSkipVerify: false,
		}
		creds := credentials.NewTLS(config)
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	conn, err := grpc.DialContext(ctx, addr, opts...)
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
