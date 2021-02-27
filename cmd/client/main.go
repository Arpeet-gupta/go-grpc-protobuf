package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/Arpeet-gupta/go-grpc-protobuf/v2/pb"
	"github.com/Arpeet-gupta/go-grpc-protobuf/v2/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dail server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	//set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//This pb.LaptopServiceClient will execute method CreateLaptop implemented by "LaptopServe struct" in laptop_server.go  and register in server/main.go using "pb.RegisterLaptopServiceServer(grpcServer, laptopServer)"
	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
	}

	log.Printf("create laptop with id: %s", res.Id)
}
