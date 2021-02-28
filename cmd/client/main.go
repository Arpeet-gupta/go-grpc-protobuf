package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"github.com/Arpeet-gupta/go-grpc-protobuf/v3/pb"
	"github.com/Arpeet-gupta/go-grpc-protobuf/v3/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createLaptop(laptopClient pb.LaptopServiceClient) {
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

func searchLaptop(laptopClient pb.LaptopServiceClient) {
	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam: &pb.Memory{
			Value: 8,
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	log.Print("search filter: ", filter)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.SearchLaptopRequest{
		Filter: filter,
	}

	stream, err := laptopClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatal("cannot search laptop: ", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		laptop := res.GetLaptop()
		log.Print("- found: ", laptop.GetId())
		log.Print("  + Name: ", laptop.GetName())
		log.Print("	 + Brand: ", laptop.GetBrand())
		log.Print("	 + CPU CORES: ", laptop.GetCpu())
		log.Print("	 + CPU MIN GHz: ", laptop.GetCpu().GetMinGhz())
		log.Print("	 + RAM: ", laptop.GetRam().GetValue(), laptop.GetRam().GetUnit())
		log.Print("	 + Price: ", laptop.GetPriceUsd(), "usd")

	}
}

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dail server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	for i := 0; i < 10; i++ {
		createLaptop(laptopClient)
	}

	searchLaptop(laptopClient)
}
