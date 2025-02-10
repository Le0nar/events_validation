package handler

import (
	"fmt"
	"net/http"

	orderevent "github.com/Le0nar/events_validation/internal/order_event"
	"github.com/gin-gonic/gin"
)

type service interface {
	SaveOrderEvent(event orderevent.OrderEvent) error
}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) SaveOrderEvent(c *gin.Context) {
	var event orderevent.OrderEvent

	// Прочитаем JSON в структуру
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Skip validation (By legends, we get those events by kafka)

	fmt.Printf("event: %v\n", event)

}

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		eventsGroup := api.Group("events")
		{
			eventsGroup.POST("/", h.SaveOrderEvent)
		}
	}

	return r
}
