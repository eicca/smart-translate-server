## Build and run
```
docker build -t translate-server .
docker run -p 8080:8080 --name tserver --rm -e GOOGLE_API_KEY=api_key translate-server
```

## Run tests
```
cp .env.example .env
edit .env
env $(cat .env | xargs) go test ./...
```
