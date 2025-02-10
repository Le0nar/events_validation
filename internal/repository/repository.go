package repository

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	orderevent "github.com/Le0nar/events_validation/internal/order_event"
)

type Repository struct {
	db driver.Conn
}

func NewRepository(db driver.Conn) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveOrderEvent(event orderevent.OrderEvent) error {
	return nil
}
