build:
	@go build -o bin/we-connect cmd/main.go

run:build
	./bin/we-connect

time:build
	time ./bin/we-connect

test:
	go test -v ./...

bench:
	go test -run none -bench . -benchtime 3s ./...