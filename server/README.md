Redis Container start command

```bash
sudo docker run -d   --name my-redis-container   -p 6379:6379   -e REDIS_PASSWORD="redis00"   -e REDIS_DB_NUMBER=0   redis:latest   redis-server --requirepass "redis00"
```