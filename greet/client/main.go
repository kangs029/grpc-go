package main

import (
	"log"

	pb "github.com/kangs029/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true // set to false if needed
	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "") // second parameter is for server name override localhost
		if err != nil {
			log.Fatalf("Error while loading CA trust certificates %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	doGreet(client)
	//doGreetManyTimes(client)
	//doLongGreet(client)
	//doGreetEveryone(client)
	//doGreetWithDeadline(client, 5*time.Second)
}
