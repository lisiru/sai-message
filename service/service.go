package service
import (
	"sai/cache"
	"sai/store"
)

type Service interface {

	Message() MessageService

}

type service struct {
	store store.Factory
	cache cache.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory,cache cache.Factory) Service {
	return &service{
		store: store,
		cache: cache,
	}
}

func (s *service) Message() MessageService  {
	return NewMessageService(s)
}

