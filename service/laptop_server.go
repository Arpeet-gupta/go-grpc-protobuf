package service

import (
	"context"
	"errors"
	"log"

	"github.com/Arpeet-gupta/go-grpc-protobuf/v2/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//LaptopServer is the server that provides laptop services
type LaptopServer struct {
	Store LaptopStore
	pb.UnimplementedLaptopServiceServer
}

//NewLaptopServer returns a new laptop server
func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

//CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// check if it's a valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop id: %v", err)
		}
		laptop.Id = id.String()
	}

	//some heavy processing
	// time.Sleep(6 * time.Second)

	if ctx.Err() == context.Canceled {
		log.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline is exceed")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	//save the laptop to in-memory store
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExits) {
			code = codes.AlreadyExists
		}
		return nil, status.Error(code, "cannot save laptop to the in-memory store")
	}
	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}
