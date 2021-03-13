package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/Arpeet-gupta/go-grpc-protobuf/v6/pb"
	"github.com/Arpeet-gupta/go-grpc-protobuf/v6/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func runRESTServer(laptopServer pb.LaptopServiceServer, listener net.Listener, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// in-process handler
	// err := pb.RegisterLaptopServiceHandlerServer(ctx, mux, laptopServer)
	err := pb.RegisterLaptopServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, dialOptions)
	if err != nil {
		return err
	}
	log.Printf("Start REST server at %s", listener.Addr().String())
	return http.Serve(listener, mux)
}

func runGRPCServer(laptopServer pb.LaptopServiceServer, listener net.Listener) error {
	grpcServer := grpc.NewServer()
	//Like we register router with server, server:= http.Server{Addr: :8080, Handler: router}, simimarly we have to register LaptopServer object with grpcServer using RegisterLaptopServiceServer() function
	//Register our service implementation with the gRPC server
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	// Like we run server.ListenAnsServer(), similarly we do  grpcServer.serve()
	log.Printf("Start GRPC server at %s", listener.Addr().String())

	return grpcServer.Serve(listener)
}

func main() {
	port := flag.Int("port", 0, "the server port")
	serverType := flag.String("type", "grpc", "type of server (grpc/rest)")
	endPoint := flag.String("endpoint", "", "gRPC endpoint")
	flag.Parse()

	//NewLaptopServer() returns object of LaptopServer Struct (LaptopServer struct implement rpc service "LaptopServiceSerer")
	//This LaptopServer struct's object  behave as router= mux.router()
	//router object holds 'HandleFunctions' for specfic PATH and Method in REST Request (router.HandleFunc("/posts", addAuthor).Methods("POST") )
	// Similarly  LaptopServer's object holds 'HandleMethods' defined by RPC service interface.
	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()
	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)

	if *serverType == "grpc" {
		err = runGRPCServer(laptopServer, listener)
	} else {
		err = runRESTServer(laptopServer, listener, *endPoint)
	}
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
