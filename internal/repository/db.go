package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func NewDB() driver.Conn {
	// Создание контекста для передачи в методы, которые требуют контекст
	ctx := context.Background()

	// Настройка подключения к ClickHouse
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"}, // Порт ClickHouse (через TCP)
		Auth: clickhouse.Auth{
			Username: "user1",   // Имя пользователя
			Password: "qwerty1", // Пароль
		},
	})
	if err != nil {
		log.Fatalf("Ошибка подключения к ClickHouse: %v", err)
	}

	// Проверка соединения с передачей контекста
	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("Не удается подключиться к ClickHouse: %v", err)
	}
	fmt.Println("Подключение к ClickHouse успешно установлено!")

	return conn
}
