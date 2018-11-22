package main_test

import (
	"context"
	"github.com/godcong/qtrago/proto"
	"github.com/godcong/qtrago/util"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func TestQtra_Start(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewQuantitativeTradingClient(conn)

	// Contact the server and print out its response.
	name := "fxxk"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.MessageInfo(ctx, &proto.MessageRequest{
		Json: string(util.Map{
			"hello": name,
		}.ToJSON()),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Json)
}
