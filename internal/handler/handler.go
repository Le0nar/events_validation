package handler

import (
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

	// Прочитаем данные из тела запроса
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Skip validation (By legends, we get those events by kafka)

	err := h.service.SaveOrderEvent(event)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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
