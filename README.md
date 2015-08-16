## Run tests
```
cp .env.example .env
edit .env
env $(cat .env | xargs) go test ./...
```
