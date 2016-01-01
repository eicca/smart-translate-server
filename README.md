# Translations and suggestions API

TODO add schema

## Prerequisites

You need docker and docker-compose installed.

## Run tests

Populate env file:
```
cp .env.example .env
edit .env
```

Start glosbe_translate service:
```
docker-compose --x-networking up -d glosbe_translate
```

Run tests:
```
docker-compose --x-networking run api go test ./...
```

## Run development cluster

```
docker-compose --x-networking up
```

## Rebuilding container

```
docker build -t translate-server .
```
