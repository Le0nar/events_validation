package main

import (
	"github.com/Le0nar/events_validation/internal/handler"
	"github.com/Le0nar/events_validation/internal/service"
)

func main() {
	service := service.NewService()
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	router.Run()
}
