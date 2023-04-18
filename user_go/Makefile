tests:
	go test ./...

build:
	go build -o dist/user_store .

build-prod:
	go build -ldflags "-s -w" -o dist/user_store .

run:
	go run .