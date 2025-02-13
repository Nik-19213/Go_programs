package main

import (
	"context"
	"fmt"
	"log"

	"github.com/codenotary/immudb/pkg/client"
	"google.golang.org/grpc/metadata"
)

func Run() error {
	client, err := client.NewImmuClient(client.DefaultOptions())
	if err != nil {
		log.Fatal(err)
		return err
	}

	ctx := context.Background()
	lr, err := client.Login(ctx, []byte("immudb"), []byte("immudb"))
	if err != nil {
		log.Fatal(err)
		return err
	}
	md := metadata.Pairs("authorization", lr.Token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	vtx, err := client.VerifiedSet(ctx, []byte("welcome"), []byte("gophers"))
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Successfully set key with transaction ID: %d\n", vtx.Id)

	ventry, err := client.VerifiedGet(ctx, []byte(`welcome`))
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Successfully retrieved and verified entry: %v\n", ventry)

	return nil
}

func main() {
	fmt.Println("Hi I am Nitin")

	if err := Run(); err != nil {
		log.Println("Error Running Application")
	}
}
