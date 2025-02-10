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
