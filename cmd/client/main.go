package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/mathyourlife/drips/proto"
)

func main() {
	fmt.Println("starting client")
	conn, err := grpc.Dial("localhost:5050",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewDripsServiceClient(conn)
	resp, err := client.Routines(context.Background(), &pb.RoutinesRequest{Name: ""})
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range resp.Routines {
		fmt.Printf("%s\n", r.Name)
	}
}
