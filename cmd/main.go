package main

import (
	"github.com/Le0nar/events_validation/internal/handler"
	"github.com/Le0nar/events_validation/internal/repository"
	"github.com/Le0nar/events_validation/internal/service"
)

func main() {
	db := repository.NewDB()

	defer db.Close()

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	router.Run()
}
