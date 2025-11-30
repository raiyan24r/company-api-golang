deps:
	go mod tidy
	go mod vendor

start:
	go run ./app/api