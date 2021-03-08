package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Arpeet-gupta/go-grpc-protobuf/v5/pb"
	"github.com/Arpeet-gupta/go-grpc-protobuf/v5/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("Start server on port %d", *port)

	//NewLaptopServer() returns object of LaptopServer Struct (LaptopServer struct implement rpc service "LaptopServiceSerer")
	//This LaptopServer struct's object  behave as router= mux.router()
	//router object holds 'HandleFunctions' for specfic PATH and Method in REST Request (router.HandleFunc("/posts", addAuthor).Methods("POST") )
	// Similarly  LaptopServer's object holds 'HandleMethods' defined by RPC service interface.
	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()
	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)
	grpcServer := grpc.NewServer()
	//Like we register router with server, server:= http.Server{Addr: :8080, Handler: router}, simimarly we have to register LaptopServer object with grpcServer using RegisterLaptopServiceServer() function
	//Register our service implementation with the gRPC server
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	// Like we run server.ListenAnsServer(), similarly we do  grpcServer.serve()
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
