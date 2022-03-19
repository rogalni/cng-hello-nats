# cng-hello-backend

Showcase of nats publish and subscribe in go.

#

## 1. How to run:

### docker-compose
```
docker build -f build/package/subscriber/docker/Dockerfile -t cng-hello-nats-subscriber .
docker build -f build/package/publisher/docker/Dockerfile -t cng-hello-nats-pulisher .

cd test/docker/cng-hello-nats

docker-compose up
```

## 2. To be done
### publish json