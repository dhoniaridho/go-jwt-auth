run:
	go run src/cmd/main.go
install:
	go mod tidy
build:
	go build -o dist/main src/cmd/main.go
test:
	go test app/main.go
