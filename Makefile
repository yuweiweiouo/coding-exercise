.PHONY: run
run: wire
	go run .

.PHONY: wire
wire: 
	wire ./...

.PHONY: test
test: wire
	go test ./...

.PHONY: build
build: wire
	go build -o dist/ .

