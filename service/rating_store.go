package service

import "sync"

// RatingStore is an interface to store laptop ratings
type RatingStore interface {
	// Add adds a new laptop score to the store and returns its rating
	Add(laptopID string, score float64) (*Rating, error)
}

// Rating contains the laptop rating information
type Rating struct {
	Count uint32
	Sum   float64
}

//InMemoryRatingStore stores laptop rating inforamtion in memory
type InMemoryRatingStore struct {
	mutex  sync.RWMutex
	rating map[string]*Rating
}

// NewInMemoryRatingStore returns new InMemoryRatingStore
func NewInMemoryRatingStore() *InMemoryRatingStore {
	return &InMemoryRatingStore{
		rating: map[string]*Rating{},
	}
}

// Add adds a new laptop score to the store and returns its rating
func (store *InMemoryRatingStore) Add(laptopID string, score float64) (*Rating, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	rating := store.rating[laptopID]
	if rating == nil {
		rating = &Rating{
			Count: 1,
			Sum:   score,
		}
	} else {
		rating.Count++
		rating.Sum += score
	}

	store.rating[laptopID] = rating
	return rating, nil
}
