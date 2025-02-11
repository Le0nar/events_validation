package repository

import (
	"context"
	"fmt"

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
	INSERT INTO order_event (event_id, order_id, user_id, event_type, event_time, order_status, total_amount)
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

// GetRecentOrderEvents получает список всех элементов из таблицы order_event за последние 15 минут
func (r *Repository) GetRecentOrderEvents() ([]orderevent.OrderEvent, error) {
	// TODO: создать табилцу с жураном проверок и брать время окончания прошлой проверки прошлой проверки

	// Формируем SQL-запрос
	query := "SELECT * FROM order_event"
	// Выполняем запрос
	rows, err := r.db.Query(context.TODO(), query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	// Преобразуем результат в слайс OrderEvent
	var events []orderevent.OrderEvent
	for rows.Next() {
		var event orderevent.OrderEvent
		if err := rows.Scan(&event.EventId, &event.OrderId, &event.UserId, &event.EventType, &event.EventTime, &event.OrderStatus, &event.TotalAmount); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		events = append(events, event)
	}

	// Проверяем на ошибки после выполнения цикла
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return events, nil
}
