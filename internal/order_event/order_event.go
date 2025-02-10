package orderevent

import (
	"time"

	"github.com/google/uuid"
)

type OrderEvent struct {
	EventId     uuid.UUID `json:"eventId" db:"event_id"`
	OrderId     uuid.UUID `json:"orderId" db:"order_id"`
	UserId      uuid.UUID `json:"userId" db:"user_id"`
	EventType   string    `json:"eventType" db:"event_type"` // Тип события (например, "order_created", "order_paid")
	EventTime   time.Time `json:"eventTime" db:"event_time"`
	OrderStatus time.Time `json:"orderStatus" db:"order_status"`
	TotalAmount time.Time `json:"totalAmount" db:"total_amount"`
}
