Redis Container start command

```bash
sudo docker run -d   --name my-redis-container   -p 6379:6379   -e REDIS_PASSWORD="redis00"   -e REDIS_DB_NUMBER=0   redis:latest   redis-server --requirepass "redis00"
```

kafka-topics.sh --create \
--topic agent-data-topic \
--bootstrap-server localhost:9092 \
--partitions 1 \
--replication-factor 1


version: '3.8'
services:
zookeeper:
image: wurstmeister/zookeeper:latest
container_name: zookeeper
environment:
ZOOKEEPER_CLIENT_PORT: 2181
ZOOKEEPER_TICK_TIME: 2000
ports:
- "2181:2181"
deploy:
resources:
limits:
memory: 256M

kafka:
image: wurstmeister/kafka:latest
container_name: kafka
depends_on:
- zookeeper
ports:
- "9092:9092"
environment:
KAFKA_BROKER_ID: 1
KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://16.171.141.95:9092
KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9092  # Bind to all interfaces
KAFKA_HEAP_OPTS: "-Xmx512M -Xms512M"       # Limit Kafka JVM memory usage to 512 MB
volumes:
- /var/run/docker.sock:/var/run/docker.sock
deploy:
resources:
limits:
memory: 512M
