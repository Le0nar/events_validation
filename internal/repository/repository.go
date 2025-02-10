package repository

import (
	"context"

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
	query := `
	INSERT INTO OrderEvent (event_id, order_id, user_id, event_type, event_time, order_status, total_amount)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	ctx := context.TODO()

	// Выполняем вставку данных в ClickHouse
	err := r.db.Exec(ctx, query, event.EventId, event.OrderId, event.UserId, event.EventType,
		event.EventTime, event.OrderStatus, event.TotalAmount)
	if err != nil {
		return err
	}

	return nil
}
