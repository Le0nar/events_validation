package service

import (
	"fmt"
	"log"
	"time"

	orderevent "github.com/Le0nar/events_validation/internal/order_event"
	"github.com/google/uuid"
)

type repository interface {
	SaveOrderEvent(orderevent.OrderEvent) error
	GetRecentOrderEvents(chan<- orderevent.OrderEvent) error
}

type Service struct {
	repo repository
}

func NewService(r repository) *Service {
	return &Service{repo: r}
}

func (s *Service) SaveOrderEvent(event orderevent.OrderEvent) error {
	return s.repo.SaveOrderEvent(event)
}

// TODO: add to config
// 600 seconds (10 min)
const tickerInterval = 600

func (s *Service) StartCheckingTicker() {
	ticker := time.NewTicker(time.Duration(time.Second * tickerInterval))
	defer ticker.Stop()

	// TODO: Можно в селекте указать еще ctx.Done(), для выключения тикера при необходимости
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		// При срабатывании таймера вызываем функцию
	// 		s.CheckOrderEvents()
	// 	}
	// }

	// TODO: вернуть цикл
	start := time.Now()

	s.CheckOrderEvents()

	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Printf("total errors: %d\n", counter)
}

// Счетчик для проверки рефакторинга
var counter int

func (s *Service) CheckOrderEvents() {
	// 1) Get list for checking

	rowsChan := make(chan orderevent.OrderEvent)

	go validateOrderEvent(rowsChan)

	err := s.repo.GetRecentOrderEvents(rowsChan)
	if err != nil {
		log.Fatal(err.Error())
	}

	// 2) Check all items in list

	// 3) log invalid item if exist (TODO: send to alert bot, if find invalid item)
}

func validateOrderEvent(rowsChan <-chan orderevent.OrderEvent) {
	for v := range rowsChan {
		if v.EventId == uuid.Nil {
			logError(v.EventId.String(), "event id is uuid nil")
		}

		if v.OrderId == uuid.Nil {
			logError(v.OrderId.String(), "order id is uuid nil")
		}

		if v.UserId == uuid.Nil {
			logError(v.UserId.String(), "user id is uuid nil")
		}

		if !isValidEventType(v.EventType) {
			logError(v.UserId.String(), "event type is invalid")
		}

		if !isValidOrderStatus(v.OrderStatus) {
			logError(v.UserId.String(), "order status is invalid")
		}
	}
}

// По легенде, ошибки пишутся отправляются на email/телегу
func logError(orderId string, errorMsg string) {
	// fmt.Printf("event order id %s error: %s \n", orderId, errorMsg)
	counter++
}

func isValidEventType(eventType string) bool {
	return eventType == orderevent.OrderCreated ||
		eventType == orderevent.OrderProcessing ||
		eventType == orderevent.OrderShipped ||
		eventType == orderevent.OrderCanceled ||
		eventType == orderevent.OrderDelivered
}

func isValidOrderStatus(orderStatus string) bool {
	return orderStatus == orderevent.StatusCreated ||
		orderStatus == orderevent.StatusProcessing ||
		orderStatus == orderevent.StatusShipped ||
		orderStatus == orderevent.StatusCanceled ||
		orderStatus == orderevent.StatusDelivered
}
