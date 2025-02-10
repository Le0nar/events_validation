# Clickhouse
1. Запуск контейнера

docker run -d --name clickhouse-server \
  -e CLICKHOUSE_DB=testdb \
  -e CLICKHOUSE_USER=user1 \
  -e CLICKHOUSE_PASSWORD=qwerty1 \
  -p 8123:8123 \
  -p 9000:9000 \
  yandex/clickhouse-server

2. Подключение к контейнеру
docker exec -it clickhouse-server clickhouse-client --user=user1 --password=qwerty1 --host=localhost

`clickhouse-server — это имя вашего контейнера ClickHouse.

clickhouse-client — клиент для взаимодействия с ClickHouse.

--user=user1 — имя пользователя, который был указан при запуске контейнера.

--password=qwerty1 — пароль для пользователя.

--host=localhost — хост, на котором работает ClickHouse (обычно это localhost, если вы подключаетесь с той же машины).`

3. Создание таблицы. После подключения к clickhouse в контейнере, нужно ввести в консоль
CREATE TABLE OrderEvent
(
    event_id UUID,               -- Уникальный идентификатор события
    order_id UUID,               -- Идентификатор заказа
    user_id UUID,                -- Идентификатор пользователя
    event_type String,           -- Тип события (например, "order_created", "order_paid")
    event_time DateTime,         -- Время события
    order_status String,         -- Статус заказа
    total_amount Float64         -- Общая сумма заказа
) 
ENGINE = MergeTree()
ORDER BY (event_time, order_id)   -- Ключ сортировки
PARTITION BY toYYYYMM(event_time);  -- Разбиение по месяцам

4. Пример вставки данных

`
INSERT INTO OrderEvent (event_id, order_id, user_id, event_type, event_time, order_status, total_amount)
VALUES
    (generateUUIDv4(), generateUUIDv4(), generateUUIDv4(), 'order_created', now(), 'pending', 100.50),
    (generateUUIDv4(), generateUUIDv4(), generateUUIDv4(), 'order_paid', now(), 'paid', 100.50),
    (generateUUIDv4(), generateUUIDv4(), generateUUIDv4(), 'order_shipped', now(), 'shipped', 100.50);
    `



# JSON
Пример тела запроса для записи
`
{
    "eventId": "43d97d38-237c-44b8-85af-edab506ef0ac",
    "orderId": "626d714f-97f5-4c18-a9e1-5630d7a7659b",
    "userId": "ff09d2ea-0a95-4008-94ad-077227d463f9",
    "eventType": "order_created",
    "createdAt": "2025-02-10T12:53:36+00:00",
    "orderStatus": "pending",
    "totalAmount": 1233
}

`
