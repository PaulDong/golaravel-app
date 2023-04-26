BINARY_NAME=bnlogicApp

build:
	@go mod vendor
	@echo "Building BNLogic..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "BNLogic built!"

run: build
	@echo "Starting BNLogic..."
	@./tmp/${BINARY_NAME} &
	@echo "BNLogic started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_compose:
	@docker-compose up -d

stop_compose:
	@docker-compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping BNLogic..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped BNLogic!"

restart: stop start