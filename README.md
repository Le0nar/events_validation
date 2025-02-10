# Clickhouse
docker run -d --name clickhouse-server \
  -e CLICKHOUSE_DB=testdb \
  -e CLICKHOUSE_USER=user1 \
  -e CLICKHOUSE_PASSWORD=qwerty1 \
  -p 8123:8123 \
  -p 9000:9000 \
  yandex/clickhouse-server





0) 
1) Create Event
2) Analyze Events every n minutes
3) Create another table for checking time 