all: build test

build: 
	@echo "Building..."
	@go build -o main cmd/api/main.go

test: 
	@echo "Testing..."
	@go test ./... -v

run: 
	@go run cmd/api/main.go

clean:
	@echo "Cleaning..."
	@rm -f main
	
docker-run: 
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi
	
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi