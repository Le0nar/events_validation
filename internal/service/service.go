package service

import orderevent "github.com/Le0nar/events_validation/internal/order_event"

type repository interface {
	SaveOrderEvent(event orderevent.OrderEvent) error
}

type Service struct {
	repo repository
}

func NewService(r repository) *Service {
	return &Service{repo: r}
}

func (s *Service) SaveOrderEvent(event orderevent.OrderEvent) error {
	return nil
}
