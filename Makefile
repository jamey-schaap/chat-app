db_data_dir = "./deployments/mysql-db/data"

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
	
docker-up: 
	@if [ $(APP_ENV) = "local" ]; then \
  		echo "Using .env file"; \
		docker_args="--env-file .env"; \
	else \
		docker_args=""; \
	fi	

	@if docker compose $(docker_args) up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up $(docker_args); \
	fi
	
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi
	
clear-db-data:
	@if [ -d $(db_data_dir) ]; then \
		rm -r $(db_data_dir); \
	fi
	
docker-clean-run: clear-db-data docker-up