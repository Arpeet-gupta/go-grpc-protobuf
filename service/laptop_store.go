package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Arpeet-gupta/go-grpc-protobuf/v2/pb"
	"github.com/jinzhu/copier"
)

//ErrAlreadyExits is returned when a record with the same ID already exist in the data store
var ErrAlreadyExits = errors.New("record already exist")

// NewInMemoryLaptopStore returns a new NewInMemoryLaptopStore object
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	// Save the laptop to the store
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// Save saves the laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if store.data[laptop.Id] != nil {
		return ErrAlreadyExits
	}

	// deep copy of the laptop object in the in-memory store
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		// return errors.New(err.Error())
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}
	store.data[other.Id] = other
	return nil
}

//Find finds a laptop by ID
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	laptop := store.data[id]

	if laptop == nil {
		return nil, nil
	}

	//deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return other, nil
}
