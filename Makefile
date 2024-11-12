# Run The Application
run:
	@go run cmd/api/main.go

# Create The DB Container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to docker compose V1"; \
		docker-compose up; \
	fi

docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to docker compose V1"; \
		docker-compose down; \
	fi

# Running Tests
test:
	@echo "Running Tests..."
	@go test -v ./...

# Clean The Binary
clean:
	@echo "Cleaning Up..."
	@rm -rf ./bin

watch:
	@if command -v air > /dev/null; then \
		air; \
		echo "Watching for changes..."; \
	else \
	        read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
	fi	

build:
	@echo "Building..."
	@go build -o bin/main cmd/api/main.go
