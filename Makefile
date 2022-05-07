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

.PHONY: docker-build
docker-build: 
	docker build -t coding-exercise . --no-cache

.PHONY: docker-run
docker-run: 
	docker stop coding-exercise || echo "try stop"
	docker rm coding-exercise || echo "try rm"
	docker run -d --name coding-exercise -p 8888:80 coding-exercise

