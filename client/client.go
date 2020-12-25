package main

import (
	"context"
	"fmt"
	"log"
	"main/proto"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting up client")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:4040", opts)
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer cc.Close()
	c := proto.NewBasicServiceClient(cc)
	//Setting Value

	fmt.Println("Setting a new value")
	value := &proto.SetRequest{Key: "1", Value: "Suman Das"}
	setRes, err := c.Set(context.Background(), value)
	if err != nil {
		log.Fatalf("Unexpected Error %v", err)
	}
	fmt.Printf("Request has been set %v", setRes)
	//Getting Value

	fmt.Println("Getting the Value")
	getVal := &proto.GetRequest{Key: "1"}
	getRes, err := c.Get(context.Background(), getVal)
	if err != nil {
		log.Fatalf("Unexpected Error %v", err)
	}
	fmt.Printf("Response :  %v", getRes)

	//Deleting Value

	fmt.Println("Deleting The value")
	delVal := &proto.DeleteRequest{Key: "1"}
	delRes, err := c.Delete(context.Background(), delVal)
	if err != nil {
		log.Fatalf("Unexpected Error %v", err)
	}
	fmt.Printf("deletion occured %v", delRes)
}
